package main_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"

	"github.com/bryack/words/adapters/acceptance"
	"github.com/bryack/words/specifications"
)

var (
	buildOnce  sync.Once
	binaryPath string
	buildError error
)

func ensureBinary() (string, error) {
	buildOnce.Do(func() {
		binPath, err := buildBinaryPath()
		if err != nil {
			buildError = err
		}
		binaryPath = binPath
	})
	return binaryPath, buildError
}

func TestWordReplacerSpecification(t *testing.T) {
	binaryPath, err := ensureBinary()
	if err != nil {
		t.Fatal(err)
	}

	// t.Cleanup выполнится только ПОСЛЕ того, как завершатся все t.Run ниже
	t.Cleanup(func() {
		err := os.RemoveAll(filepath.Dir(binaryPath))
		if err != nil {
			t.Logf("Warning: failed to remove temp directory: %v", err)
		}
	})

	t.Run("should replace words in text", func(t *testing.T) {
		tempDir := t.TempDir()
		dataFile := filepath.Join(tempDir, "test.jsonl")
		err := createTestJSONLFile(t, dataFile)
		if err != nil {
			t.Fatal(err)
		}
		driver := &acceptance.Driver{
			BinaryPath: binaryPath,
			DataFile:   dataFile,
			TempDir:    tempDir,
		}
		specifications.WordReplacerSpecification(t, driver)
	})
}

func createTestJSONLFile(t *testing.T, dataFile string) error {
	t.Helper()
	data := []byte("{\"senses\":[{\"links\":[[\"counterfeit\",\"counterfeit\"],[\"imitation\",\"imitation\"],[\"fake\",\"fake\"],[\"forgery\",\"forgery\"]],\"synonyms\":[{\"word\":\"фальша́к\"},{\"word\":\"фальши́вка\"}],\"glosses\":[\"counterfeit,imitation;fake;forgery\"],\"categories\":[\"Pageswith1entry\",\"Pageswithentries\",\"Russian3-syllablewords\",\"Russianentrieswithincorrectlanguageheader\",\"Russianfemininenouns\",\"Russianinanimatenouns\",\"Russianlemmas\",\"Russianlinkswithredundantaltparameters\",\"Russianlinkswithredundantwikilinks\",\"Russiannouns\",\"Russiannounswithaccentpatterna\",\"Russiannounswithreduciblestem\",\"Russiantermssuffixedwith-ка\",\"RussiantermswithIPApronunciation\",\"Russianvelar-stemfeminine-formaccent-anouns\",\"Russianvelar-stemfeminine-formnouns\",\"ru:Crime\"]}],\"pos\":\"noun\",\"head_templates\":[{\"name\":\"ru-noun+\",\"args\":{\"1\":\"подде́лка\",\"2\":\"*\"},\"expansion\":\"подде́лка•(poddélka)finan(genitiveподде́лки,nominativepluralподде́лки,genitivepluralподде́лок)\"}],\"forms\":[{\"form\":\"подде́лка\",\"tags\":[\"canonical\",\"feminine\",\"inanimate\"]},{\"form\":\"poddélka\",\"tags\":[\"romanization\"]},{\"form\":\"подде́лки\",\"tags\":[\"genitive\"]},{\"form\":\"подде́лки\",\"tags\":[\"nominative\",\"plural\"]},{\"form\":\"подде́лок\",\"tags\":[\"genitive\",\"plural\"]},{\"form\":\"no-table-tags\",\"source\":\"declension\",\"tags\":[\"table-tags\"]},{\"form\":\"ru-noun-table\",\"source\":\"declension\",\"tags\":[\"inflection-template\"]},{\"form\":\"velar-stem\",\"source\":\"declension\",\"tags\":[\"class\"]},{\"form\":\"accent-a\",\"source\":\"declension\",\"tags\":[\"class\"]},{\"form\":\"подде́лка\",\"tags\":[\"nominative\",\"singular\"],\"source\":\"declension\",\"roman\":\"poddélka\"},{\"form\":\"подде́лки\",\"tags\":[\"nominative\",\"plural\"],\"source\":\"declension\",\"roman\":\"poddélki\"},{\"form\":\"подде́лки\",\"tags\":[\"genitive\",\"singular\"],\"source\":\"declension\",\"roman\":\"poddélki\"},{\"form\":\"подде́лок\",\"tags\":[\"genitive\",\"plural\"],\"source\":\"declension\",\"roman\":\"poddélok\"},{\"form\":\"подде́лке\",\"tags\":[\"dative\",\"singular\"],\"source\":\"declension\",\"roman\":\"poddélke\"},{\"form\":\"подде́лкам\",\"tags\":[\"dative\",\"plural\"],\"source\":\"declension\",\"roman\":\"poddélkam\"},{\"form\":\"подде́лку\",\"tags\":[\"accusative\",\"singular\"],\"source\":\"declension\",\"roman\":\"poddélku\"},{\"form\":\"подде́лки\",\"tags\":[\"accusative\",\"plural\"],\"source\":\"declension\",\"roman\":\"poddélki\"},{\"form\":\"подде́лкой\",\"tags\":[\"instrumental\",\"singular\"],\"source\":\"declension\",\"roman\":\"poddélkoj\"},{\"form\":\"подде́лкою\",\"tags\":[\"instrumental\",\"singular\"],\"source\":\"declension\",\"roman\":\"poddélkoju\"},{\"form\":\"подде́лками\",\"tags\":[\"instrumental\",\"plural\"],\"source\":\"declension\",\"roman\":\"poddélkami\"},{\"form\":\"подде́лке\",\"tags\":[\"prepositional\",\"singular\"],\"source\":\"declension\",\"roman\":\"poddélke\"},{\"form\":\"подде́лках\",\"tags\":[\"plural\",\"prepositional\"],\"source\":\"declension\",\"roman\":\"poddélkax\"}],\"etymology_text\":\"подде́лать(poddélatʹ,“tocounterfeit,tofake”)+-ка(-ka)\",\"etymology_templates\":[{\"name\":\"af\",\"args\":{\"1\":\"ru\",\"2\":\"подде́лать\",\"t1\":\"tocounterfeit,tofake\",\"3\":\"-ка\"},\"expansion\":\"подде́лать(poddélatʹ,“tocounterfeit,tofake”)+-ка(-ka)\"}],\"sounds\":[{\"ipa\":\"[pɐˈdʲːeɫkə]\"},{\"audio\":\"Ru-подделка.ogg\",\"ogg_url\":\"https://upload.wikimedia.org/wikipedia/commons/5/51/Ru-%D0%BF%D0%BE%D0%B4%D0%B4%D0%B5%D0%BB%D0%BA%D0%B0.ogg\",\"mp3_url\":\"https://upload.wikimedia.org/wikipedia/commons/transcoded/5/51/Ru-%D0%BF%D0%BE%D0%B4%D0%B4%D0%B5%D0%BB%D0%BA%D0%B0.ogg/Ru-%D0%BF%D0%BE%D0%B4%D0%B4%D0%B5%D0%BB%D0%BA%D0%B0.ogg.mp3\"}],\"word\":\"подделка\",\"lang\":\"Russian\",\"lang_code\":\"ru\"}")
	err := os.WriteFile(dataFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", data, err)
	}
	return nil
}

func buildBinaryPath() (string, error) {
	tempDir, err := os.MkdirTemp("", "test-binary-*")
	if err != nil {
		return "", fmt.Errorf("failed to make temp directory: %v", err)
	}
	binName := "temp-binary"
	binPath := filepath.Join(tempDir, binName)

	build := exec.Command("go", "build", "-cover", "-o", binPath, ".")
	build.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("GOCOVERDIR"))
	var stderr bytes.Buffer
	build.Stderr = &stderr

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Go build Error: \n%q\n", stderr.String())
		return "", fmt.Errorf("cannot build tool %s: %v", binName, err)
	}
	return binPath, nil
}
