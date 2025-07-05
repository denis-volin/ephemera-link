document.addEventListener('DOMContentLoaded', function() {
  // Fix link text and href
  const linkElement = document.getElementById('copy');
  const currentUrl = window.location.href;
  const link = linkElement.textContent;
  linkElement.href = currentUrl + link;
  linkElement.textContent = currentUrl + link;

  // Show expire time in human-readable format
  const expireElement = document.getElementById('expire');
  const seconds = expireElement.textContent;
  const now = new Date();
  const expireDate = new Date(now.getTime() + seconds * 1000);

  const formatter = new Intl.RelativeTimeFormat(["en", "ru"]);
  const diffInSeconds = Math.floor((expireDate - now) / 1000);

  if (diffInSeconds < 60) {
    expireElement.textContent = formatter.format(diffInSeconds, 'second');
    return
  }

  const diffInMinutes = Math.floor(diffInSeconds / 60);
  if (diffInMinutes < 60) {
    expireElement.textContent = formatter.format(diffInMinutes, 'minute');
    return
  }

  const diffInHours = Math.floor(diffInMinutes / 60);
  if (diffInHours < 24) {
    expireElement.textContent = formatter.format(diffInHours, 'hour');
    return
  }

  const diffInDays = Math.floor(diffInHours / 24);
  expireElement.textContent = formatter.format(diffInDays, 'day');
});
