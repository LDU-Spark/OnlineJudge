document.getElementById('registerForm').addEventListener('submit', function(e) {
    e.preventDefault();
    
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirmPassword').value;
    
    if (password !== confirmPassword) {
        document.getElementById('message').textContent = 'Passwords do not match';
        return;
    }
    
    fetch('/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            username: username,
            password: password
        }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            window.location.href = '/logic';
        } else {
            document.getElementById('message').textContent = data.message;
        }
    })
    .catch((error) => {
        console.error('Error:', error);
        document.getElementById('message').textContent = 'An error occurred. Please try again.';
    });
});