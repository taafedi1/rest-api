# Aufgaben für die REST API 

## Aufgabe 6: Todo REST Backend I
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