
# kanban-board-cli

**For Linux (x64):**
```bash
curl -L https://github.com/Shivam583-hue/kanban-board-cli-/releases/latest/download/kanban-linux-amd64 -o kanban
chmod +x kanban
sudo mv kanban /usr/local/bin/
kanban
```

**For macOS (Intel):**
```bash
curl -L https://github.com/Shivam583-hue/kanban-board-cli-/releases/latest/download/kanban-darwin-amd64 -o kanban
chmod +x kanban
sudo mv kanban /usr/local/bin/
kanban
```

**For macOS (M1/M2):**
```bash
curl -L https://github.com/Shivam583-hue/kanban-board-cli-/releases/latest/download/kanban-darwin-arm64 -o kanban
chmod +x kanban
sudo mv kanban /usr/local/bin/
kanban
```

**For Windows (using PowerShell):**
```powershell
Invoke-WebRequest -Uri https://github.com/Shivam583-hue/kanban-board-cli-/releases/latest/download/kanban-windows-amd64.exe -OutFile kanban.exe
# Move to a directory in your PATH or run from current directory
.\kanban.exe
```

After installation, you can run the tool by simply typing:
```bash
kanban
```

Controls:
- `←/h`, `→/l`: Navigate between columns
- `↑/k`, `↓/j`: Navigate tasks
- `n`: Create new task
- `enter`: Move task to next column
- `q`: Quit


---
## To contribute
 ```bash
git clone https://github.com/Shivam583-hue/kanban-board-cli-
cd kanban-board-cli-
go mod tidy
go run main.go
```
---
## Screenshots
![image](https://github.com/user-attachments/assets/fb2f52c6-2933-4a27-8f70-d19bcc5756fd)
