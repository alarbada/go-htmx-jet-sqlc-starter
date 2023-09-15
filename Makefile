run:
	air

css:
	bunx tailwindcss -i ./templates/main.css -o ./public/main.css --watch

sql:
	sqlc generate

db:
	sqlite3 ./tmp/app.sqlite

fmt:
	go fmt ./...
	bunx prettier --write  ./templates/*.tmpl ./templates/**/*.tmpl

