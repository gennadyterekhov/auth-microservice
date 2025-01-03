go test -v -coverpkg=./... -coverprofile=artefacts/coverage/coverage.out -covermode=count ./... > /dev/null

#  ignore generated
cat artefacts/coverage/coverage.out | grep -v ".pb." > artefacts/coverage/coverage.temp.out
cat artefacts/coverage/coverage.temp.out | grep -v "/swagger/" > artefacts/coverage/coverage.out
cat artefacts/coverage/coverage.out | grep -v "/repository_error_mock" > artefacts/coverage/coverage.temp.out
cat artefacts/coverage/coverage.temp.out | grep -v "/repository_mock" > artefacts/coverage/coverage.out
rm artefacts/coverage/coverage.temp.out

# must have go 1.23 to use
covreport -i artefacts/coverage/coverage.out -o artefacts/coverage/cover.html -cutlines 80,40
echo "http://localhost:8080/artefacts/coverage/cover.html"
python3 -m http.server 8080