docker-build:
	echo "Clippr_Img set to: $(CLIPPR_IMG)" 
	docker build -f Dockerfile . -t ${CLIPPR_IMG}

docker-push:
	docker push ${CLIPPR_IMG}

# Package all binaries using goreleaser
package: $(GORELEASER)
	$(GORELEASER) --snapshot --skip-publish --rm-dist

