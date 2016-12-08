# Path Truncation

La page du [challenge](https://www.root-me.org/fr/Challenges/Web-Serveur/Path-Truncation)

## Analyse

On constate qu’une LFI (Local File Inclusion) est possible via la variable " ?page="

En PHP (v < 5.3), les strings ne peuvent supporter qu'une chaîne de longueur maximum égal à 4096 caractères. Si on dépasse la longueur, PHP va tout simplement tronqué la chaîne.

En PHP, quand on ajoute un fichier en donnant son chemin. Le php va former du code en réduisant le nom des fichiers, ce qui nous permet de donner des noms de fichiers très long en ajoutant `/.` autant de fois que l'on veut. Par exemple si on donne la variable `"/file/././."` d'un fichier que l'on veut inclure, php va le traduire en `include("/file");`.

La taille du chemin doit rester impaire pour que l'attaque marche. On ajoute toujours deux caractères `/.`, pour avoir une taille impair, on va ajouter un symbole inutile, ici `/`, après le nom de la page, `?page=admin.html//././././././.`, au lieu d'avoir `?page=admin.html/././././././.`.

## Attaque

On construit l'URL suivante, qu'on entre ensuite dans un navigateur internet.

```sh
page=$(python -c 'print "admin.html/"+"/."*2050') 
url="http://chall/index.php?page=$page"
echo $url
```
