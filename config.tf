provider "google" {
  credentials = file("service-account.json")

  project = "photostest-1564886711039"
  region  = "us-central1"
  zone    = "us-central1-c"
}

# google_cloudfunctions_function.default:
# imported with the command:
#  terraform import google_cloudfunctions_function.default photostest-1564886711039/us-central1/Tao
#  terraform show
resource "google_cloudfunctions_function" "default" {
    available_memory_mb   = 256
    entry_point           = "Tao"
    environment_variables = {}
    https_trigger_url     = "https://us-central1-photostest-1564886711039.cloudfunctions.net/Tao"
    id                    = "photostest-1564886711039/us-central1/Tao"
    labels                = {
        "deployment-tool" = "cli-gcloud"
    }
    max_instances         = 0
    name                  = "Tao"
    project               = "photostest-1564886711039"
    region                = "us-central1"
    runtime               = "go111"
    service_account_email = "photostest-1564886711039@appspot.gserviceaccount.com"
    timeout               = 60
    trigger_http          = true

    timeouts {}
}
