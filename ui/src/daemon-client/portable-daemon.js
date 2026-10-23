import { spawn } from "child_process";
import path from "path";
import { GetPortableDaemonDir } from "@/helpers/main_platform";

// PORTABLE MODE (Windows): start the daemon located next to the UI (elevated, shows UAC prompt).
// Returns true if the start was initiated.

export function IsPortableDaemonAvailable() {
  return GetPortableDaemonDir() != null;
}

const psQuote = (s) => "'" + String(s).replace(/'/g, "''") + "'";

export function TryStartPortableDaemon() {
  const dir = GetPortableDaemonDir();
  if (!dir) return false;
  const exe = path.join(dir, "IVPN Service.exe");
  try {
    const cmd = `Start-Process -FilePath ${psQuote(exe)} -WorkingDirectory ${psQuote(dir)} -WindowStyle Hidden -Verb RunAs`;
    const child = spawn(
      "powershell.exe",
      ["-NoProfile", "-NonInteractive", "-WindowStyle", "Hidden", "-Command", cmd],
      { detached: true, stdio: "ignore", windowsHide: true }
    );
    child.on("error", (err) => console.error("Failed to start portable daemon:", err));
    child.unref();
    console.log("Starting portable daemon:", exe);
    return true;
  } catch (e) {
    console.error("Failed to start portable daemon:", e);
    return false;
  }
}
