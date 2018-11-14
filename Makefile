DEV_APPSERVER ?= dev_appserver.py

export GCLOUD_PROJECT ?= local-storage-security

export ROOT_PATH = .

export PORT ?= 8081

# `default` is listed first so that the default service is deployed first.
# This satisfies the Google App Engine condition that a default service must be
# deployed first to a new project.
DIRS    = default analytics app attacker
TARGETS = deploy

$(TARGETS:%=\%.%):
	$(MAKE) -C $* $(@:$*.%=%)

$(TARGETS):
	$(MAKE) $(DIRS:%=%.$@)

start:
	$(DEV_APPSERVER) \
		-A $(GCLOUD_PROJECT) \
		--enable_host_checking=false \
		--port=$(PORT) \
		$(ROOT_PATH)/analytics/app.yaml \
		$(ROOT_PATH)/app/app.yaml \
		$(ROOT_PATH)/attacker/app.yaml

.PHONY: $(TARGETS) start
