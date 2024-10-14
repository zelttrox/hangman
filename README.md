# hangman

Hangman est un jeu du pendu sur Terminal créé en Go.

> Créé en utilisant [VS Code](https://code.visualstudio.com/)

**Auteurs :**
- [Romain](https://github.com/Velapsis)
- [Enzo](https://github.com/zelttrox)
- [Charleen](https://github.com/CharLuney)

-------------------------------------

**INSTALLATION POUR UBUNTU**

_à partir du terminal Ubuntu_

Installation de Go :
```
sudo apt-get -y install golang-go
```

Installation de Arcatia :
```
git clone https://github.com/zelttrox/hangman.git
```

Lancement du jeu :
```
go run main.go -list [Dictionnaire]
```
Il existe deux dictionnaires :
- Jeux ("games")
- League of Legends ("lol")

Le jeu possède aussi un mode difficile (Hardcore).
Lancement en mode difficile :
```
go run main.go -hard
```
