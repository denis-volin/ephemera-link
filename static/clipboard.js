document.getElementById('copyBtn').addEventListener('click', async () => {
  const linkElement = document.getElementById('copy');
  const linkText = linkElement.textContent.trim(); // Get the link text

  try {
    // Modern clipboard API (most browsers)
    await navigator.clipboard.writeText(linkText);

    // Optional: Show feedback
    const originalText = document.getElementById('copyBtn').textContent;
    document.getElementById('copyBtn').textContent = 'Copied!';
    setTimeout(() => {
      document.getElementById('copyBtn').textContent = originalText;
    }, 2000);

  } catch (err) {
    // Fallback for older browsers
    const textarea = document.createElement('textarea');
    textarea.value = linkText;
    textarea.style.position = 'fixed'; // Avoid scrolling to bottom
    document.body.appendChild(textarea);
    textarea.select();

    try {
      document.execCommand('copy'); // Legacy method
      alert('Link copied to clipboard!'); // Fallback feedback
    } catch (err) {
      console.error('Failed to copy:', err);
      prompt('Press Ctrl+C to copy:', linkText); // Last resort
    } finally {
      document.body.removeChild(textarea);
    }
  }
});
