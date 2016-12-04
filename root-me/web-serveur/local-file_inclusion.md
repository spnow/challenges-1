# Local File Inclusion

La page du [challenge](http://challenge01.root-me.org/web-serveur/ch16/)

## Principe

On va exploiter une vulnérabilité LFI (Local File Inclusion), qui permet à un utilisateur d’inclure des fichiers locaux (appartenant au serveur) à partir d’une URL.

## Attaque

En analysant la page on apercoit:

* Présence d'un répertoire _admin_ grâce au lien à droite dans le menu.
* Utilisation des variables _files_ et _file_ dans l'URL qui nous permette de sélectionner des répertoires et fichiers.

On commencer donc par tester si on peut aller dans des répertoires interdit: `http://challenge01.root-me.org/web-serveur/ch16/?files=../`. On trouve alors une liste:

* text.gif
* admin
* files
* index.php

On a remarqué la présence d'un dossier _admin_, qu'on va sélectionner: `http://challenge01.root-me.org/web-serveur/ch16/?files=../admin`. On trouve alors une liste:

* index.php

On clique sur ce fichier, ou alors on utilise la variable _f_ dans l'url: `http://challenge01.root-me.org/web-serveur/ch16/?files=../admin&f=index.php`

En regardant dans le contenue du fichier, on trouve le flag du challenge.
