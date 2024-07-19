document.addEventListener('DOMContentLoaded', function() {
    const button = document.getElementById('change-text-btn');
    const paragraph = document.getElementById('demo-text');

    button.addEventListener('click', function() {
        paragraph.textContent = 'The text has been changed!';
    });
});
