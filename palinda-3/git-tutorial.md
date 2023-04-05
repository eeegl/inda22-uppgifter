# Git Bootcamp

## Innehåll

Den här texten går igenom hur vi skapar ett Git-repo och uppdaterar förändringar. För en kort lista med de mest grunläggande git-kommandona, se [lathunden](git-lathund.md).

* [Grunder](#grunder)
    * [Initialisera Git-repo](#initialisera-git-repo)
    * [Registrera ändringar](#registrera-ändringar)
* [Mer om ändringar och hur man ångrar sig](#mer-om-ändringar-och-hur-man-ångrar-sig)
    * [`README.md` och kompakt vy](#readmemd-och-kompakt-vy)
    * [Ångra och gå tillbaka till senaste commit](#ångra-och-gå-tillbaka-till-senaste-commit)
    * [Tillägg till senaste commit](#tillägg-till-senaste-commit)

## Grunder

### Initialisera Git-repo

Öppna terminalen och navigera till dokument-mappen (eller annan valfri mapp).

```bash
cd ~
cd Documents/
```

Skapa en ny mapp som heter `"hellogit"` och gå in i mappen.

```bash
mkdir hellogit
cd hellogit/
```

För att initialisera en mapp till ett git-repo (förkortning av git-*repository*) använder vi:

```bash
git init .
``` 

*Glöm inte punkten efter `git init`! Den refererar till den nuvarande mappen (`hellogit` i vårt fall).*

Detta skapar en gömd mapp (`.git`) som håller koll på ändringarna i git-repot. Du borde även se något som liknar den här outputen:

```bash
tips: Använder "master" som namn för den inledande grenen. Detta förvalda grennamn
tips: kan ändras i framtiden. För att välja vilket namn som ska användas på
tips: den inledande grenen i alla nya arkiv, och dölja denna varning, kör du:
tips:
tips: 	git config --global init.defaultBranch <namn>
tips:
tips: Namn som ofta används istället för "master" är "main", "trunk" och
tips: "development". Den nyskapade grenen kan ges nytt namn med kommandot:
tips:
tips: 	git branch -m <namn>
Initierade tomt Git-arkiv i /Users/orn/Documents/hellogit/.git/
```

Vi kan nu kontrollera att Git fungerar genom att kolla Git-statusen med `git status`. Detta borde ge info om att vi är på master-grenen och inte har lagt till något ännu:

```bash
På grenen master

Inga incheckningar ännu

inget att checka in (skapa/kopiera filer och spåra med "git add")
```

Bra jobbat, du har initaliserad Git-repot!

### Registrera ändringar

För att göra en förändring i repot, låt oss skapa filen `HelloGit.java`:

```
touch HelloGit.java
``` 

Redigera filen genom att lägga till följande kod:

```java
class HelloGit {
    public static void main(String[] args) {
        System.out.println("Hello, Git!");
    }
}
```

Kör vi `git status` igen kommer vi se något som liknar denna output:

```bash
På grenen master

Inga incheckningar ännu

Ospårade filer:
  (använd "git add <fil>..." för att ta med i det som ska checkas in)
	HelloGit.java

inget köat för incheckning, men ospårade filer finns (spåra med "git add")
```

Git har alltså noterat att filen existerar, men spårar den inte ännu. För att börja spåra filen använder vi `git add`:

```
git add HelloGit.java
``` 

En ny `git status` borde ge något som liknar detta:

```bash
På grenen master

Inga incheckningar ännu

Ändringar att checka in:
  (använd "git rm --cached <fil>..." för att ta bort från kö)
	ny fil:        HelloGit.java
```

Nu är alltså `HelloGit.java` redo för incheckning (engelska: *commit*). Vi checkar in våra ändringar med `git commit`:

```bash
git commit
```

Detta öppnar upp en textredigerare där du kan lägga till ett meddelande som förklarar vilka ändringar som har gjorts i incheckningen:

```

# Ange incheckningsmeddelandet för dina ändringar. Rader som inleds
# med "#" kommer ignoreras, och ett tomt meddelande avbryter incheckningen.
#
# På grenen master
#
# Första incheckning
#
# Ändringar att checka in:
#	ny fil:        HelloGit.java
#
```

Skriv ett kort och kärnfullt meddelande som förklarar vad du har gjort, till exempel "Första commit: skapar HelloGit.java".

En ny `git status` ger följande meddelande:

```bash
På grenen master
inget att checka in, arbetskatalogen ren
```

Detta betyder att vi inte har gjort några ändringar sedan senaste incheckningen, vilket inte är så konstigt eftersom vi nyss gjorde den!

Men hur kan vi se vår incheckning då? För det använder vi `git log`:

```bash
commit f26b8ecf210f46f7a1ed188bfdf35fc9e49cf0b9 (HEAD -> master)
Author: ornse <ornse@kth.se>
Date:   Wed Apr 5 10:40:27 2023 +0200

    Första commit: skapar HelloGit.java
```

Nu ser vi vår incheckning med vårt meddelande! Vi ser också att den har blivit tilldelad ett *inchecknings-ID*: `f26b8ecf210f46f7a1ed188bfdf35fc9e49cf0b9`. Detta ID använvänds för att referera till olika incheckningar.

Snyggt jobbat! Ta gärna en titt och se så att du förstår vad alla kommandon gör och hur du använder dem. Känner du dig säker kan du gå vidare till nästa del. :)

## Mer om ändringar och hur man ångrar sig

### `README.md` och kompakt vy
När vi gör vårt repo offentligt vill vi att folk som besöker det ska kunna få en överblick av innehållet. Vi gör det med en [`README.md`](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-readmes) som sammanfattar vad syftet med vår repo är. Skapa en README-fil med en kort beskrivning av ditt repo i [Markdown-format](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet), till exempel:

```
# Mitt Github-repo

Det här är ett övningsrepo som innehåller filen [HelloGit.java](/HelloGit.java) som printar `Hello, Git!`.
```

Vi lägger till och checkar in vår README. För att slippa gå in i textredigeraren och skriva vårt incheckningsmeddelande kan vi lägga till `m`-flaggan till `git commit`, följt av meddelandet inom citattecken:

```bash
git add README.md
git commit -m "Lägger till README.md"
```

Nu visar `git log` båda våra incheckningar:

```bash
commit 2ed794cae84b0c4917d4b99ac54f94283f93791b (HEAD -> master)
Author: ornse <ornse@kth.se>
Date:   Wed Apr 5 11:45:22 2023 +0200

    Lägger till README.md

commit f26b8ecf210f46f7a1ed188bfdf35fc9e49cf0b9
Author: ornse <ornse@kth.se>
Date:   Wed Apr 5 10:40:27 2023 +0200

    Första commit: skapar HelloGit.java
```

> *Tips: för en mer kompakt vy kan vi använda `git log --oneline`.*

### Ångra och gå tillbaka till senaste commit

Vi är glada att vi har lärt oss så mycket om Git redan, så vi ändrar print-satsen i `HelloGit.java` så att den istället skriver ut `HELLO, GIT!!!`.

För att se ändringen kan vi använda `git diff`, som ger något som liknar följande output:

```bash
diff --git a/HelloGit.java b/HelloGit.java
index 26a27a2..55cfdc5 100644
--- a/HelloGit.java
+++ b/HelloGit.java
@@ -1,6 +1,6 @@
 class HelloGit {
     public static void main(String[] args) {
-        System.out.println("Hello, Git!");
+        System.out.println("HELLO, GIT!!!");
     }
 }
```

Under headern för `main()` ser vi en rad med minustecken och en med plustecken framför. Raden med minustecken som visar vad som tagits bort och den med plustecken vad som har lagts till sedan senaste ändring.

Ändringen ser bra ut så vi lägger till den med `git add`!...

...men när vi tänker efter var det nog faktiskt bättre som det var innan, så vi vill ändra tillbaka till hur det var vid vår senaste incheckning. Hur gör vi det? Med `git reset --hard`! 

> Tips: `git reset --hard` rensar alla ändringar och går tillbaka till senaste incheckningen, vilket kan vara skönt om du vet att inte vill spara de senaste ändringarna. Men var försiktigt så att du inte råkar radera fler ändringar än du hade planerat, kommandot går inte att ångra!

### Tillägg till senaste commit

Antag att vi vill lägga till en extra rad i vår `HelloGit.java` så att den ser ut såhär:

```java
class HelloGit {
    public static void main(String[] args) {
        System.out.println("Hello, Git!");
        System.out.println("Goodbye, Git!"); // Ny rad
    }
}
```

Ändringen är väldigt liten och det känns inte riktigt som att den rättfärdigar en helt ny incheckning. Istället kan vi göra ett tillägg till vår förra commit:

```bash
git add HelloGit.java
git commit --amend
```

### Vidare läsning

För mer info om Git, kolla in de här länkarna:

* [Git Book](https://git-scm.com/book/en/v2)
* [Atlassians Git tutorial](https://www.atlassian.com/git/tutorials/setting-up-a-repository)

Lycka till!
