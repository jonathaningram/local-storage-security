// Simulate setting the token once the user successfully logs in.

const token = Math.random().toString();

localStorage.setItem("token", token);

document.getElementById("user_token").textContent = token;
