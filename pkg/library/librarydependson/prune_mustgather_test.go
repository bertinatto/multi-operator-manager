package librarydependson

import (
	"context"
	"github.com/deads2k/multi-operator-manager/pkg/applyconfiguration"
	"github.com/deads2k/multi-operator-manager/pkg/library/libraryapplyconfiguration"
	"os"
	"path"
	"sigs.k8s.io/yaml"
	"strings"
	"testing"
)

func TestGetRequiredResourcesFromMustGather(t *testing.T) {
	// this exists to make it easy to generate output for the unit tests
	writeActualContent := true

	content, err := os.ReadDir("test-data")
	if err != nil {
		t.Fatal(err)
	}

	for _, currTestDir := range content {
		ctx := context.Background()
		testName := currTestDir.Name()
		t.Run(testName, func(t *testing.T) {
			testPertinentResources := path.Join("test-data", currTestDir.Name(), "pertinent-resources.yaml")
			pertinentResourcesBytes, err := os.ReadFile(testPertinentResources)
			if err != nil {
				t.Fatal(err)
			}
			pertinentResources := &PertinentResources{}
			if err := yaml.Unmarshal(pertinentResourcesBytes, &pertinentResources); err != nil {
				t.Fatal(err)
			}

			mustGatherDirPath := path.Join("test-data", currTestDir.Name(), "input-dir")
			expectedDirPath := path.Join("test-data", currTestDir.Name(), "expected-output")

			if writeActualContent {
				err = WriteRequiredResourcesFromMustGather(ctx, pertinentResources, mustGatherDirPath, expectedDirPath)
				if err != nil {
					t.Fatal(err)
				}
				return
			}

			actualPertinentResources, err := GetRequiredResourcesFromMustGather(ctx, pertinentResources, mustGatherDirPath)
			if err != nil {
				t.Fatal(err)
			}

			expectedPertinentResources, err := libraryapplyconfiguration.LenientResourcesFromDirRecursive(expectedDirPath)
			if err != nil {
				t.Fatal(err)
			}

			differences := applyconfiguration.EquivalentResources("pruned", expectedPertinentResources, actualPertinentResources)
			if len(differences) > 0 {
				t.Log(strings.Join(differences, "\n"))
				t.Errorf("expected results mismatch %d times with actual results", len(differences))
			}
		})
	}
}
