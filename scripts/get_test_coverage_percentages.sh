go test -v -coverpkg=./... -coverprofile=artefacts/coverage/coverage.out -covermode=count ./... > /dev/null

cat artefacts/coverage/coverage.out | grep -v ".pb." > artefacts/coverage/coverage.temp.out
cat artefacts/coverage/coverage.temp.out | grep -v "/swagger/" > artefacts/coverage/coverage.out
cat artefacts/coverage/coverage.out | grep -v "/repository_error_mock" > artefacts/coverage/coverage.temp.out
cat artefacts/coverage/coverage.temp.out | grep -v "/repository_mock" > artefacts/coverage/coverage.out
rm artefacts/coverage/coverage.temp.out

go tool cover -func artefacts/coverage/coverage.out > artefacts/coverage/coverage.txt.out
rev artefacts/coverage/coverage.txt.out | tail -n 1 | cut -c -5 | rev