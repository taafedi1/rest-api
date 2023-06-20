# Aufgaben für die REST API 

## Aufgabe 6: Todo REST Backend I [v1.0.0]
Dieser Aufgabe beiliegend bekommst du durch den Dozenten ein bereits lauffähiges Go Todo REST Ba-
ckend, wo ein REST Backend für eine Todo Liste implementiert ist. Dieses Projekt basiert auf dem Http-
Router als leichtgewichtiger, hochleistungsfähiger HTTP-Request-Router. Zudem hat diese Applikation
bereits eine Datenpersistenz mittels CSV-Datei implementiert und ist nach dem MVC-Pattern implemen-
tiert, wobei keine View-Komponente hier zum Einsatz kommt, da es nur ein Backend ist.
Folgende Ressourcen sind bereits implementiert:

| No. | HTTP Verb | Path      | Expect (JSON) | Returns (JSON) | HTTP Status | Description  |
|-----|-----------|-----------|---------------|----------------|-------------|--------------|
|1    |GET        |/          |Nothing        |Welcome string  |200(success) |Welcome string|
|2    |GET        |/todos     |Nothing        |an aaray with todo enteries  |200(success) |Get a list of todos|
|3    |GET        |/todos/:id |Nothing        |The todo specified ID  |200(success) or 404 (not found) |Get todo by ID|
|4    |POST       |/todos     |A todo entry   |The new todo entry  |201(created) 400 (bad request) |create new todo|

Optional:
- Erstelle auf GitHub in deinem Account ein Repository für dieses Todo REST Backend Projekt
und checke das initiale Projekt dort ein. Dann kannst du im fortlaufenden an deinem Projekt mit-
tels Versionsverwaltungssystem weiterarbeiten.

Mache dich vertraut mit dem Quellcode und lasse das Programm laufen. Fürs Ausprobieren der einzel-
nen Endpunkte verwenden wir besser eine API-Plattform als einen Browser, da der Browser für nicht
GET-Endpunkte schwierig zu verwenden wäre. Eine bewährte API-Plattform, welche auch kostenlos ver-
wendet werden kann, ist Postman. Installiere bei dir lokal Postman, welches du herunterladen kannst un-
ter (Hinweis: Auch ohne Registrierung und Login verwendbar):
https://www.postman.com/
Verwende nun Postman fürs Ausprobieren der einzelnen Endpunkte.
![alt text](https://github.com/taafedi1/rest-api/blob/main/images/rest.png)

Hier ein Beispiel für den 4. Endpunkt:
![alt text](https://github.com/taafedi1/rest-api/blob/main/images/rest_bsp1.png)
Anmerkung: Hier musst du entsprechend einen Body konfigurieren und dort das neu zu erstellende Todo
JSON Objekt drin aufführen.

## Aufgabe 7: Todo REST Backend II [v2.0.0]
Nun sollen die Endpunkte erweitert werden. Neuerdings soll mittels Put eine bestehende Todo Ressource
verändert werden können. Dieser neue Endpunkt ist in der folgenden Tabelle auf Zeile 5 aufgeführt:
| No. | HTTP Verb | Path      | Expect (JSON) | Returns (JSON) | HTTP Status | Description  |
|-----|-----------|-----------|---------------|----------------|-------------|--------------|
|1    |GET        |/          |Nothing        |Welcome string  |200(success) |Welcome string|
|2    |GET        |/todos     |Nothing        |an aaray with todo enteries  |200(success) |Get a list of todos|
|3    |GET        |/todos/:id |Nothing        |The todo specified ID  |200(success) or 404 (not found) |Get todo by ID|
|4    |POST       |/todos     |A todo entry   |The new todo entry  |201(created) 400 (bad request) |create new todo|
|5 |PUT |/todos:id | A todo entry |The updated todo entry| 200(success) 400 (bad request) 404 (not found)| Update todo by ID|

Implementiere nun diesen neu geforderten Endpunkt gemäss der Spezifikation in der Tabelle in Zeile 5.
Teste den Endpunkt wiederum dann mittels Postman auf Korrektheit.

## Aufgabe 8: Todo REST Backend III
Nun sollen die Endpunkte erweitert werden. Neuerdings soll mittels Delete eine bestehende Todo Res-
source entfernt werden können. Dieser neue Endpunkt ist in der folgenden Tabelle auf Zeile 6 aufgeführt:
| No. | HTTP Verb | Path      | Expect (JSON) | Returns (JSON) | HTTP Status | Description  |
|-----|-----------|-----------|---------------|----------------|-------------|--------------|
|1    |GET        |/          |Nothing        |Welcome string  |200(success) |Welcome string|
|2    |GET        |/todos     |Nothing        |an aaray with todo enteries  |200(success) |Get a list of todos|
|3    |GET        |/todos/:id |Nothing        |The todo specified ID  |200(success) or 404 (not found) |Get todo by ID|
|4    |POST       |/todos     |A todo entry   |The new todo entry  |201(created) 400 (bad request) |create new todo|
|5 |PUT |/todos:id | A todo entry |The updated todo entry| 200(success) 400 (bad request) 404 (net found)| Update todo by ID|
|6 |DELETE |/todos:id | Nothing |Nothing| 200(success) 404 (not found)| Delete todo by ID|

Implementiere nun diesen neu geforderten Endpunkt gemäss der Spezifikation in der Tabelle in Zeile 6.
Durchs Löschen von einzelnen Todo’s müssen wir die Reihenfolge der IDs gegebenenfalls noch korrigie-
ren, damit keine Lücken entstehen. Beachte bei der Implementierung, dass wenn nicht das letzte Ele-
ment entfernt wird, dann bei allen nach dem löschenden Element kommenden Todos die ID angepasst
werden muss (d.h. jeweils dekrementiert werden). Die IDs sollen mit dem Index 0 starten und dann auf-
steigend bis n-1 hochgezählt werden, wobei n hier die Anzahl Todos meint.
Teste den Endpunkt wiederum dann mittels Postman auf Korrektheit.

## Aufgabe 9: Todo REST Backend IIII [Optional]
Mache dir nun Gedanken, wie du das Todo REST Backend noch erweiterten könntest. Du könntest z.B.
noch weitere Endpunkte mitaufnehmen oder z.B. das Todo-Datenmodell erweitern. Hier kannst du selbst
bestimmen was aus deiner Sicht noch Sinn macht und das dann auch implementieren.
Beschreibe dann in deinem README.md was du für Erweiterungen gemacht hast, sodass es dokumen-
tiert ist und auch später einfach nachvollziehbar ist.