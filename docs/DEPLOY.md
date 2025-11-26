Perfect! I've set up everything you need to host your Go project on dev-manager.signalvision.ai. Here's what has been configured:

Setup Complete

1. Go Installation
- Go 1.23.3 (ARM64) installed at /usr/local/go
- Add to your PATH: export PATH=$PATH:/usr/local/go/bin

2. Directory Structure
/home/devmanager/
├── htdocs/dev-manager.signalvision.ai/  (your Go app goes here)
└── logs/
    ├── nginx/
    └── app/

3. Supervisor Configuration
- File: /etc/supervisor/conf.d/devmanager.conf
- Will run your app from: /home/devmanager/htdocs/dev-manager.signalvision.ai/app
- Port: 8081

4. Nginx Configuration
- File: /etc/nginx/sites-enabled/dev-manager.signalvision.ai.conf
- Proxies HTTPS traffic to your Go app on port 8081

Next Steps to Deploy Your Go Project:

1. Upload your Go project:
# On your local machine, upload to the server:
scp -r your-go-project/* ubuntu@your-server:/home/devmanager/htdocs/dev-manager.signalvision.ai/

# Or use git:
cd /home/devmanager/htdocs/dev-manager.signalvision.ai
git clone your-repo-url .

2. Build your Go application:
cd /home/devmanager/htdocs/dev-manager.signalvision.ai
/usr/local/go/bin/go build -o app .

3. Set up SSL certificate:
You'll need to obtain SSL certificates for dev-manager.signalvision.ai. Place them at:
- /etc/nginx/ssl-certificates/dev-manager.signalvision.ai.crt
- /etc/nginx/ssl-certificates/dev-manager.signalvision.ai.key

4. Configure your Go app:
Make sure your Go application:
- Listens on port 8081 (or read from PORT environment variable)
- Example:
package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Go!")
    })

    fmt.Printf("Server starting on port %s\n", port)
    http.ListenAndServe(":"+port, nil)
}

5. Start the services:
sudo supervisorctl reread
sudo supervisorctl update
sudo supervisorctl start devmanager
sudo nginx -t && sudo systemctl reload nginx

6. Check status:
sudo supervisorctl status devmanager
curl http://localhost:8081

Once you have your SSL certificates in place and your Go app deployed, your application will be accessible at https://dev-manager.signalvision.ai!

Would you like help deploying a specific Go project, or do you need assistance with the SSL certificate setup?

∴ Thought for 1s (ctrl+o to show thinking)