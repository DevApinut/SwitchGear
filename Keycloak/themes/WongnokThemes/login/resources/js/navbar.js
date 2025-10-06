// Custom Navigation Bar for Keycloak Login
document.addEventListener('DOMContentLoaded', function() {
    console.log('🚀 Loading navbar...');
    
    // Create navigation bar
    const navbar = document.createElement('div');
    navbar.className = 'custom-navbar';
    navbar.innerHTML = `
        <a href="https://localhost:8443" class="nav-home">
            <span class="nav-icon">🏠</span>
            <span class="nav-text">Home</span>
        </a>           
    `;
    
    // Insert at the beginning of body
    document.body.insertBefore(navbar, document.body.firstChild);
    console.log('✅ Navbar added to page');
    
    // Add click event listener for debugging
    const homeLink = navbar.querySelector('.nav-home');
    homeLink.addEventListener('click', function(e) {
        console.log('🖱️ Home link clicked!');
        console.log('🔗 Going to:', e.target.href || e.currentTarget.href);
    });
});


