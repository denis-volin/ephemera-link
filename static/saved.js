document.addEventListener('DOMContentLoaded', function() {
  // Fix link text and href
  const linkElement = document.getElementById('copy');
  const currentUrl = window.location.href;
  const link = linkElement.textContent;
  linkElement.href = currentUrl + link;
  linkElement.textContent = currentUrl + link;
});
