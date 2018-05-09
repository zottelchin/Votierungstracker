# Votierungstracker

Diese Webapp entwickel ich hauptsächlich für mich selber. Ich möchte gerne eine schnelle, minimale Möglichkeit meine Vortierungen/Vorträge und Punkte in Kursen zu speichern.

Die App soll über https://vote.zottelchin.de erreichbar sein.

## Roadmap

| User Story          | Beschreibung                                                                                                                                                       |
| ==========          | ============                                                                                                                                                       |
| U01 Server & Login  | Die Webapp kann über die URL aufgerufen werden, Accounts werden per Secret verwaltet.                                                                              |
| U02 Veranstaltungen | Auf dem Startbildschirm sieht man die Veranstaltungen, die der jeweilige Nutzer angelegt hat.                                                                      |
| U03 Votierungen     | Auf der Seite einer Veranstaltung sieht man den Namen der Veranstaltung, die Punkte sowie die möglichen Punkte. Einträge können hinzugefügt und bearbeitet werden. |
| U04 Fortschritt     | Pro Kurs kann die zu erreichende Prozent-/Punktzahl eingetragen werden; es wird der Fortschritt angezeigt. **Milestone 1.0**                                       |
| U05 Sicherheit      | User-Accounts gebunden an E-Mail-Adresse, Passwortwiederherstellung, Unterstützung von U2F, Security-Header, optionales HTTPS mit Let's Encrypt. **Milestone 1.5** |

## Struktur

**Frontend:** [Vue.js](https://vuejs.org/) und [Milligram](https://milligram.io/)
⬍ REST+JSON
**Backend**: [Golang](https://golang.org/) mit [Gin](https://gin-gonic.github.io/gin/) als Webframework und [SQLite](https://www.sqlite.org/) für die Datenspeicherung

## Einrichtung

1. Benötigte Software: [Go](https://golang.org/) & [node.js](https://nodejs.org/)
2. run Build-Tool installieren: `go get github.com/moqmar/run`
3. dep Dependency-Tool installieren: `go get github.com/golang/dep/cmd/dep`
4. Code herunterladen: `go get github.com/zottelchin/Votierungstracker`
5. Wechsel ins Code-Verzeichnis: `cd $GOPATH/src/github.com/zottelchin/Votierungstracker`
5. Frontend-Dependencies installieren: `cd frontend && { npm i; cd ..; }`
6. Backend-Dependencies installieren: `dep ensure`
7. Programm bauen, starten und bei Dateiänderungen automatisch aktualisieren: `run`
