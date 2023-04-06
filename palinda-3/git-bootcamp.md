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
* [Städat repo och förgrening](#städat-repo-och-förgrening)
    * [Strunta i filer listade i `.gitignore`](#strunta-i-filer-listade-i-gitignore)
    * [Förgrening](#förgrening)
    * [Merge conflict](#merge-conflict)
* [Vidare läsning](#vidare-läsning)

> *Sektionen "Städat repo och förgrening" är skriven av Mathias Grindsäter, som även har bidragit till resten av texten.*

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

## Städat repo och förgrening

### Strunta i filer listade i `.gitignore`
I ditt repo, skapa en ny fil `GitImportance.java` och kopiera in följande kod till filen.
```java
import java.util.*;

public class GitImportance{ 
    public static void main(String[] args) {
        StringBuilder s = new StringBuilder();
        s.append("Git allows for version control, making it difficult to track changes in code.\n");
        s.append("With Git, only one developer can collaborate on a codebase ");
        s.append("without interfering with each other's work.\n");
        
        System.out.println(s.toString());
    }
}
```

Kompilera och kör `GitImportance.java` och se till att den faktiskt skriver ut
innehållet till terminalen.

Vår nya fil är dock ospårad. Som vi nu har lärt oss kan vi se detta genom att
i terminalen skriva 

```bash
git status
``` 

Där bör vi se något som ser ut ungefär såhär:

```bash
On branch master
Untracked files:
  (use "git add <file>..." to include in what will be committed)
	GitImportance.class
	GitImportance.java

nothing added to commit but untracked files present (use "git add" to track)
```

Vill vi verkligen behöva ta hänsyn till `.class`-filen? Vi kan skapa en fil
`.gitignore`. Då kan vi säga till Git att ignorera vissa filer.
Skapa filen genom att skriva 
``` bash
touch .gitignore
```

Öppna filen och skriv in följande:
```back
### JAVA ###
*.class
```

Kör `git status` i terminalen. Det borde se ut ungefär såhär:
```bash
On branch master
Untracked files:
  (use "git add <file>..." to include in what will be committed)
	.gitignore
	GitImportance.java

nothing added to commit but untracked files present (use "git add" to track)
```

Vår .class-fil är borta! Git ignorerar nu alla .class-filer. Vi har 
fått ett nytt tillskott `.gitignore`, men denna vill vi ha med oss i fortsättningen!

Skriv nu i terminalen. 

```bash
git add .gitignore GitImportance.java

git commit -m "lägger till .gitignore and GitImportance.java"
```

Kör sedan `git status`. I terminalen ska det nu stå:
```bash
On branch master
nothing to commit, working tree clean
```

Vi bör nu se 3 commits om vi i terminalen skriver:
```bash
git log --oneline
```

>Tips: Du kan använda en .gitignore-fil från
>era inda-repon. Dessa är fyllda av
>filtyper vi helst vill ignorera.

### Förgrening
Kompilera och kör programmet `GitImportance.java` igen.
Läser vi noggrant så noterar vi att det som skrivs ut inte är helt korrekt.
Låt oss skapa en gren, där vi kan arbeta på en lösning utan att det påverkar
vår ursprungliga kod.

Skapa en branch som heter `fixaText` genom att skriva följande i terminalen:

`git branch fixaText`.

För att se att vår gren faktiskt har skapats. Skriv följande i terminalen:

`git branch`

Vi bör få något som ser ut såhär:
```bash
  fixaText
* master
```

Asterisken betyder att vi för närvarande är på master-grenen. För att byta till vår nya gren `fixaTest`
skriver vi:

`git switch fixaText`

Kör sedan 

`git branch` 

Dubbelkolla att asterisken befinner sig på vår nya gren:
```bash
* fixaText
  master
```

Låt oss nu fixa texten! I filen `GitImportance.java`, byt ut `"difficult"` mot `"easy"`
och `"only one developer"` mot `"multiple developers"`.

Efter att ha utfört ändringen, kompilera och kör programmet. Kontrollera
att utskriften är korrekt. Skriv sedan 

`git status` 

i terminalen. Som svar bör vi få:
```bash
On branch fixaText
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	modified:   GitImportance.java

no changes added to commit (use "git add" and/or "git commit -a")
```

Vi ser att Git har noterat förändringen i `GitImportance.java`.
För att se exakt vad som ändrats kan vi köra `git diff`.

Vi kan nu förbereda för incheckning (stage:a) och checka in vår rättning. Vi utför då som tidigare
följande två steg i terminalen.
```bash
git add GitImportance.java
git commit -m "Rättade texten i GitImportance.java"
```

Vi har nu arbetat i vår gren `fixaText` och vet att vi har rättat texten där. 
Men hur ser det ut i `master`-grenen? Vi byter till `master`-grenen i terminalen genom att skriva

`git switch master`

Testa sedan att kompilera och köra `GitImportance.java`.

Vi ser att texten här inte är rättad, vilket inte är så konstigt, eftersom vår
rättade version finns i en annan branch `fixText`.

Nu vill vi sammanslå (merge:a) vår rättade text i `fixaText`-grenen till vår `master`-gren. Detta gör vi genom att i terminalen skriva:

```brash
git merge -m "Sammanslår fixaText med master" fixaText
```
Terminalen bör svara med:
```bash
Updating 14ccad4..8e69362
Fast-forward (no commit created; -m option ignored)
 GitImportance.java | 4 ++--
 1 file changed, 2 insertions(+), 2 deletions(-)
 ```

Kompilera och kör nu programmet i master-branchen och kontrollera att du nu får rätt utskrift.

Eftersom vi är färdiga med `fixaText`-grenen så kan vi radera den
genom att skriva följande i terminalen:
```bash
git branch -d fixaText
```
Kontrollera nu att `fixaText`-grenen är raderad genom att skriva `git branch`.

### Merge conflict
Någon har sagt att variabeln `s` i filen `GitImportance.java`
är otydlig. Vi bestämmer oss därför för att döpa om den 
till `text`. Vi skapar därför en ny gren genom att i terminalen skriva:

```bash
git switch -c variabelByte
```

Notera att vi här både skapar och byter till vår nya gren i ett kommando.

Ändra filen `GitImportance.java` så att det ser ut såhär:

```java
import java.util.*;

public class GitImportance{ 
    public static void main(String[] args) {
        StringBuilder text = new StringBuilder();
        text.append("Git allows for version control, making it easy to track changes in code.\n");
        text.append("With Git, multiple developers can collaborate on a codebase ");
        text.append("without interfering with each other's work.\n");
        
        System.out.println(text.toString());
    }
}
```

Kompilera och kör programmet för att kontrollera att det funkar.

Vi förbereder för incheckning (stage:ar) och checkar in genom att i terminalen skriva:

```bash
git add GitImportance.java

git commit -m "döpte om variabeln s till text"
```

Vi har nu döpt om vår variabel och byter därför tillbaka till master-grenen:

```bash
git switch master
```

Vi kommer nu på att vi vill lägga till en rad med information om Git i vårt program.
Vi är något ouppmärksamma och öppnar därför filen `GitImportance.java` i master-grenen -
det vill säga den gren vi nu befinner oss i - och lägger till raden i filen `GitImportance.java`:

```java
s.append("Using Git can increase productivity by streamlining workflows and reducing the time spent on manual tasks.\n");
```

Kompilera och kör programmet för att se att den nya rader skrivs ut.

Sedan förbereder vi för incheckning och checkar in genom att skriva följande i terminalen:

```bash
git add GitImportance.java

git commit -m "la till en rad i GitImportance.java"
```

Vi kommer nu på att vi inte har slagit ihop (merge:at) våra två grenar.

Vi slår därför ihop grenen `variabelByte` till `master`-grenen genom att
skriva följande i terminalen:

```bash
git merge variabelByte
```

Git säger dock:
```bash
CONFLICT (content): Merge conflict in GitImportance.java
Automatic merge failed; fix conflicts and then commit the result.
```

Eftersom vi har ändringar i båda våra grenar `master` och `variabelByte` så får vi
en `merge conflict`. Om du öppnar filen i exempelvis VSC, eller din vanliga texthanterare, så
bör du se kod från båda grenar. Lös merge-konflikten så att vi både har variabeln `s` döpt till `text` och
att vi får med vår nya rad. När du har löst detta kan du gå tillbaka till terminalen och skriva:

```bash
git add GitImportance.java

git commit -m "la till en rad i GitImportance.java och döpte om variabeln s till text"
```

Kontrollera nu att filen innehåller det den ska. Kompilera sedan och kör programmet. Du kan sedan radera grenen 
`variabelByte`.

## Vidare läsning

För mer info om Git, kolla in de här länkarna:

* [Git Book](https://git-scm.com/book/en/v2)
* [Atlassians Git tutorial](https://www.atlassian.com/git/tutorials/setting-up-a-repository)

Lycka till!