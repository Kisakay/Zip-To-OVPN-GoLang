package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func combineFilesToOvpnFolder(folderPath string) {
	// Vérifier si le dossier existe
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		fmt.Println("Le dossier spécifié n'existe pas.")
		return
	}

	// Vérifier si tous les fichiers requis existent dans le dossier
	requiredFiles := []string{"ca.crt", "client.crt", "client.key", "openvpn.ovpn"}
	missingFiles := []string{}
	for _, file := range requiredFiles {
		if _, err := os.Stat(filepath.Join(folderPath, file)); os.IsNotExist(err) {
			missingFiles = append(missingFiles, file)
		}
	}
	if len(missingFiles) > 0 {
		fmt.Println("Les fichiers suivants sont manquants dans le dossier :")
		fmt.Println(strings.Join(missingFiles, "\n"))
		return
	}

	// Lire le contenu des fichiers
	fileContents := make(map[string]string)
	for _, file := range requiredFiles {
		content, err := ioutil.ReadFile(filepath.Join(folderPath, file))
		if err != nil {
			fmt.Printf("Erreur lors de la lecture du fichier %s : %v\n", file, err)
			return
		}
		fileContents[file] = string(content)
	}

	// Supprimer les lignes mentionnant les fichiers de certificats et de clés
	keysToRemove := []string{"ca ca.crt", "cert client.crt", "key client.key"}
	for _, key := range keysToRemove {
		fileContents["openvpn.ovpn"] = strings.Replace(fileContents["openvpn.ovpn"], key+"\n", "", -1)
	}

	// Créer le fichier combined.ovpn
	combinedOvpnPath := filepath.Join(folderPath, "combined.ovpn")
	combinedOvpnFile, err := os.Create(combinedOvpnPath)
	if err != nil {
		fmt.Printf("Erreur lors de la création du fichier combined.ovpn : %v\n", err)
		return
	}
	defer combinedOvpnFile.Close()

	combinedOvpnFile.WriteString(fileContents["openvpn.ovpn"])
	combinedOvpnFile.WriteString("<ca>\n")
	combinedOvpnFile.WriteString(strings.Replace(fileContents["ca.crt"], "ca ca.crt", "ez", -1))
	combinedOvpnFile.WriteString("\n</ca>\n\n<cert>\n")
	combinedOvpnFile.WriteString(strings.Replace(fileContents["client.crt"], "cert client.crt", "", -1))
	combinedOvpnFile.WriteString("\n</cert>\n\n<key>\n")
	combinedOvpnFile.WriteString(strings.Replace(fileContents["client.key"], "key client.key", "", -1))
	combinedOvpnFile.WriteString("\n</key>\n\n")

	fmt.Printf("Le fichier combined.ovpn a été créé avec succès dans le dossier %s.\n", folderPath)
}

func main() {
	fmt.Print("Entrez le chemin du dossier contenant les fichiers (ex: /chemin/vers/le/dossier) : ")
	var folderPath string
	fmt.Scanln(&folderPath)
	combineFilesToOvpnFolder(folderPath)
}
