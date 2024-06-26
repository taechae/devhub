package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/taechae/devhub/pkg/attestation"
	"github.com/taechae/devhub/pkg/types"
)

const breeds = "{\"current_page\":1,\"data\":[{\"breed\":\"Abyssinian\",\"country\":\"Ethiopia\",\"origin\":\"Natural\\/Standard\",\"coat\":\"Short\",\"pattern\":\"Ticked\"},{\"breed\":\"Aegean\",\"country\":\"Greece\",\"origin\":\"Natural\\/Standard\",\"coat\":\"Semi-long\",\"pattern\":\"Bi- or tri-colored\"},{\"breed\":\"American Curl\",\"country\":\"United States\",\"origin\":\"Mutation\",\"coat\":\"Short\\/Long\",\"pattern\":\"All\"},{\"breed\":\"American Bobtail\",\"country\":\"United States\",\"origin\":\"Mutation\",\"coat\":\"Short\\/Long\",\"pattern\":\"All\"},{\"breed\":\"American Shorthair\",\"country\":\"United States\",\"origin\":\"Natural\",\"coat\":\"Short\",\"pattern\":\"All but colorpoint\"},{\"breed\":\"American Wirehair\",\"country\":\"United States\",\"origin\":\"Mutation\",\"coat\":\"Rex\",\"pattern\":\"All but colorpoint\"},{\"breed\":\"Arabian Mau\",\"country\":\"Arabian Peninsula\",\"origin\":\"Natural\",\"coat\":\"Short\",\"pattern\":\"\"},{\"breed\":\"Australian Mist\",\"country\":\"Australia\",\"origin\":\"Crossbreed\",\"coat\":\"Short\",\"pattern\":\"Spotted and Classic tabby\"},{\"breed\":\"Asian\",\"country\":\"developed in the United Kingdom (founding stock from Asia)\",\"origin\":\"\",\"coat\":\"Short\",\"pattern\":\"Evenly solid\"},{\"breed\":\"Asian Semi-longhair\",\"country\":\"United Kingdom\",\"origin\":\"Crossbreed\",\"coat\":\"Semi-long\",\"pattern\":\"Solid\"},{\"breed\":\"Balinese\",\"country\":\"developed in the United States (founding stock from Thailand)\",\"origin\":\"Crossbreed\",\"coat\":\"Long\",\"pattern\":\"Colorpoint\"},{\"breed\":\"Bambino\",\"country\":\"United States\",\"origin\":\"Crossbreed\",\"coat\":\"Hairless\\/Furry down\",\"pattern\":\"\"},{\"breed\":\"Bengal\",\"country\":\"developed in the United States (founding stock from Asia)\",\"origin\":\"Hybrid\",\"coat\":\"Short\",\"pattern\":\"Spotted\\/Marbled\"},{\"breed\":\"Birman\",\"country\":\"developed in France (founding stock from Burma)\",\"origin\":\"Natural\",\"coat\":\"Semi Long\",\"pattern\":\"Colorpoint\"},{\"breed\":\"Bombay\",\"country\":\"developed in the United States (founding stock from Asia)\",\"origin\":\"Crossbred\",\"coat\":\"Short\",\"pattern\":\"Solid\"},{\"breed\":\"Brazilian Shorthair\",\"country\":\"Brazil\",\"origin\":\"Natural\",\"coat\":\"Short\",\"pattern\":\"All\"},{\"breed\":\"British Semi-longhair\",\"country\":\"United Kingdom\",\"origin\":\"\",\"coat\":\"Medium\",\"pattern\":\"All\"},{\"breed\":\"British Shorthair\",\"country\":\"United Kingdom\",\"origin\":\"Natural\",\"coat\":\"Short\",\"pattern\":\"All\"},{\"breed\":\"British Longhair\",\"country\":\"United Kingdom\",\"origin\":\"\",\"coat\":\"Long\",\"pattern\":\"\"},{\"breed\":\"Burmese\",\"country\":\"Burma and Thailand\",\"origin\":\"Natural\",\"coat\":\"Short\",\"pattern\":\"Solid\"},{\"breed\":\"Burmilla\",\"country\":\"United Kingdom\",\"origin\":\"Crossbreed\",\"coat\":\"Short\\/Long\",\"pattern\":\"\"},{\"breed\":\"California Spangled\",\"country\":\"United States\",\"origin\":\"Crossbreed\",\"coat\":\"Short\",\"pattern\":\"Spotted\"},{\"breed\":\"Chantilly-Tiffany\",\"country\":\"United States\",\"origin\":\"\",\"coat\":\"\",\"pattern\":\"\"},{\"breed\":\"Chartreux\",\"country\":\"France\",\"origin\":\"Natural\",\"coat\":\"Short\",\"pattern\":\"Solid\"},{\"breed\":\"Chausie\",\"country\":\"France\",\"origin\":\"Hybrid\",\"coat\":\"Short\",\"pattern\":\"Ticked\"}],\"first_page_url\":\"https:\\/\\/catfact.ninja\\/breeds?page=1\",\"from\":1,\"last_page\":4,\"last_page_url\":\"https:\\/\\/catfact.ninja\\/breeds?page=4\",\"links\":[{\"url\":null,\"label\":\"Previous\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/breeds?page=1\",\"label\":\"1\",\"active\":true},{\"url\":\"https:\\/\\/catfact.ninja\\/breeds?page=2\",\"label\":\"2\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/breeds?page=3\",\"label\":\"3\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/breeds?page=4\",\"label\":\"4\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/breeds?page=2\",\"label\":\"Next\",\"active\":false}],\"next_page_url\":\"https:\\/\\/catfact.ninja\\/breeds?page=2\",\"path\":\"https:\\/\\/catfact.ninja\\/breeds\",\"per_page\":25,\"prev_page_url\":null,\"to\":25,\"total\":98}"
const fact = "{\"fact\":\"MMMohammed loved cats and reportedly his favorite cat, Muezza, was a tabby. Legend says that tabby cats have an \\u201cM\\u201d for Mohammed on top of their heads because Mohammad would often rest his hand on the cat\\u2019s head.\",\"length\":210}"
const facts = "{\"current_page\":1,\"data\":[{\"fact\":\"Unlike dogs, cats do not have a sweet tooth. Scientists believe this is due to a mutation in a key taste receptor.\",\"length\":114},{\"fact\":\"When a cat chases its prey, it keeps its head level. Dogs and humans bob their heads up and down.\",\"length\":97},{\"fact\":\"The technical term for a cat\\u2019s hairball is a \\u201cbezoar.\\u201d\",\"length\":54},{\"fact\":\"A group of cats is called a \\u201cclowder.\\u201d\",\"length\":38},{\"fact\":\"A cat can\\u2019t climb head first down a tree because every claw on a cat\\u2019s paw points the same way. To get down from a tree, a cat must back down.\",\"length\":142},{\"fact\":\"Cats make about 100 different sounds. Dogs make only about 10.\",\"length\":62},{\"fact\":\"Every year, nearly four million cats are eaten in Asia.\",\"length\":55},{\"fact\":\"There are more than 500 million domestic cats in the world, with approximately 40 recognized breeds.\",\"length\":100},{\"fact\":\"Approximately 24 cat skins can make a coat.\",\"length\":43},{\"fact\":\"While it is commonly thought that the ancient Egyptians were the first to domesticate cats, the oldest known pet cat was recently found in a 9,500-year-old grave on the Mediterranean island of Cyprus. This grave predates early Egyptian art depicting cats by 4,000 years or more.\",\"length\":278}],\"first_page_url\":\"https:\\/\\/catfact.ninja\\/facts?page=1\",\"from\":1,\"last_page\":34,\"last_page_url\":\"https:\\/\\/catfact.ninja\\/facts?page=34\",\"links\":[{\"url\":null,\"label\":\"Previous\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=1\",\"label\":\"1\",\"active\":true},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=2\",\"label\":\"2\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=3\",\"label\":\"3\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=4\",\"label\":\"4\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=5\",\"label\":\"5\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=6\",\"label\":\"6\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=7\",\"label\":\"7\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=8\",\"label\":\"8\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=9\",\"label\":\"9\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=10\",\"label\":\"10\",\"active\":false},{\"url\":null,\"label\":\"...\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=33\",\"label\":\"33\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=34\",\"label\":\"34\",\"active\":false},{\"url\":\"https:\\/\\/catfact.ninja\\/facts?page=2\",\"label\":\"Next\",\"active\":false}],\"next_page_url\":\"https:\\/\\/catfact.ninja\\/facts?page=2\",\"path\":\"https:\\/\\/catfact.ninja\\/facts\",\"per_page\":10,\"prev_page_url\":null,\"to\":10,\"total\":332}"

func main() {
	// Test vulnerability detection.
	exec.Command(os.Args[0], os.Args[1:]...)

	log.Print("starting server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/breeds", breedHandler)
	http.HandleFunc("/fact", factHandler)
	http.HandleFunc("/facts", factsHandler)

	http.HandleFunc("/runtimes", runHandler)
	http.HandleFunc("/runtime-artifacts", runArtifactsHandler)
	http.HandleFunc("/vulnerabilities", topVulHandler)
	http.HandleFunc("/vulnerable-artifacts", vulArtifactsHandler)
	http.HandleFunc("/packages", pkgHandler)
	http.HandleFunc("/builds", buildHandler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func breedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, breeds)
}

func factHandler(w http.ResponseWriter, r *http.Request) {
	// ***
	var bodyBytes []byte
	var err error
	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Body reading error: %v", err)
			return
		}
		defer r.Body.Close()
	}

	fmt.Fprintf(w, "Headers: %+v\n", r.Header)

	if len(bodyBytes) > 0 {
		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
			fmt.Fprintf(w, "JSON parse error: %v", err)
			return
		}
		fmt.Fprintf(w, "%s", string(prettyJSON.Bytes()))
	} else {
		fmt.Fprintf(w, "Body: No Body Supplied\n")
	}

	fmt.Fprintf(w, "Query: %+s\n", r.URL.RawQuery)
	// ***

	fmt.Fprintf(w, fact)
}

func factsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, facts)
}

func findArtifact(artifacts []string, needle string) bool {
	for _, a := range artifacts {
		if a == needle {
			return true
		}
	}
	return false
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	opt := &types.RunOptions{
		Project:  "s3c100",
		Location: "us-central1",
	}

	artifacts := strings.Split(r.URL.Query().Get("artifacts"), ",")

	out, err := attestation.GetRunRevisions(r.Context(), opt)
	if err != nil {
		log.Fatal(err)
	}

	for i, r := range out {
		if findArtifact(artifacts, r.ArtifactURI) {
			out[i].Match = true
		}
	}

	b, _ := json.Marshal(struct {
		Data []attestation.RunRevision
	}{
		Data: out,
	})
	fmt.Fprintf(w, string(b))
}

func vulHandler(w http.ResponseWriter, r *http.Request) {
	cve := r.URL.Query().Get("vulnerability")

	opt := &types.VulnOptions{
		Project: "s3c100",
		Cve:     cve,
	}

	out, err := attestation.GetVulnerabilities(r.Context(), opt)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(struct {
		Data []attestation.Vulnerability
	}{
		Data: out,
	})
	fmt.Fprintf(w, string(b))
}

func topVulHandler(w http.ResponseWriter, r *http.Request) {
	// ***
	var bodyBytes []byte
	var err error
	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Body reading error: %v", err)
			return
		}
		defer r.Body.Close()
	}

	fmt.Fprintf(w, "Headers: %+v\n", r.Header)

	if len(bodyBytes) > 0 {
		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
			fmt.Fprintf(w, "JSON parse error: %v", err)
			return
		}
		fmt.Fprintf(w, "%s", string(prettyJSON.Bytes()))
	} else {
		fmt.Fprintf(w, "Body: No Body Supplied\n")
	}

	fmt.Fprintf(w, "Query: %+s\n", r.URL.RawQuery)
	// ***

	return

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	opt := &types.VulnOptions{
		Project: "s3c100",
		Limit:   limit,
	}

	out, err := attestation.GetTopVulnerabilities(r.Context(), opt)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(struct {
		Data []attestation.TopVulnerabilities
	}{
		Data: out,
	})
	fmt.Fprintf(w, string(b))
}

func runArtifactsHandler(w http.ResponseWriter, r *http.Request) {
	opt := &types.RunOptions{
		Project:  "s3c100",
		Location: "us-central1",
	}

	out, err := attestation.GetRunRevisions(r.Context(), opt)
	if err != nil {
		log.Fatal(err)
	}

	var artifacts []string
	for _, r := range out {
		artifacts = append(artifacts, r.ArtifactURI)
	}

	b, _ := json.Marshal(artifacts)
	fmt.Fprintf(w, string(b))
}

func vulArtifactsHandler(w http.ResponseWriter, r *http.Request) {
	cve := r.URL.Query().Get("vulnerability")
	artifact := r.URL.Query().Get("artifact")

	opt := &types.VulnOptions{
		Project:     "s3c100",
		Cve:         cve,
		ArtifactURI: artifact,
	}

	out, err := attestation.GetVulnerabilities(r.Context(), opt)
	if err != nil {
		log.Fatal(err)
	}

	var yes bool
	if len(out) > 0 {
		yes = true
	}

	b, _ := json.Marshal(yes)
	fmt.Fprintf(w, string(b))

	/*

		b, _ := json.Marshal(struct {
			Data []attestation.Vulnerability
		}{
			Data: out,
		})
		fmt.Fprintf(w, string(b))
	*/
}

func pkgHandler(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	p := r.URL.Query().Get("package")

	opt := &types.PkgOptions{
		Project: "s3c100",
		Limit:   limit,
		Package: p,
	}

	out, err := attestation.GetPackages(r.Context(), opt)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(struct {
		Data []attestation.Package
	}{
		Data: out,
	})
	fmt.Fprintf(w, string(b))
}

func buildHandler(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	opt := &types.BuildOptions{
		Project: "s3c100",
		Limit:   limit,
	}

	out, err := attestation.GetBuilds(r.Context(), opt)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(struct {
		Data []attestation.Build
	}{
		Data: out,
	})
	fmt.Fprintf(w, string(b))
}
