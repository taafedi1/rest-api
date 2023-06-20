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
|2    |GET        |/todos     |Nothing        |an aaray with\  |200(success) |Welcome string|
|     |           |           |               |todo enteries   |             |              |
|3    |GET        |/todos/:id |Nothing        |Welcome string  |200(success) |Welcome string|
|4    |POST       |/todos     |A todo entry   |Welcome string  |200(success) |Welcome string|

--------------------------------------------------------
Spalte links       Spalte rechts   Spalte mittig
---------------- ---------------- ----------------------
In dieser Zelle\   Hier nicht!     In dieser Zelle wird 
werden zwei\                       kein Umbruch gewagt, 
Umbrüche gewagt.                   aber mehr Text 
                                   geschrieben.      

Zeile 2          Zeile 2          Zeile 2

Hier wird eine\    Hier  nicht!    Hier auch nicht!
Zeile gebrochen.
--------------------------------------------------------