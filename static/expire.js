function relativeTime(seconds, locale) {
  const now = new Date();
  const expireDate = new Date(now.getTime() + seconds * 1000);

  const formatter = new Intl.RelativeTimeFormat(locale);
  const diffInSeconds = Math.floor((expireDate - now) / 1000);

  if (diffInSeconds < 60) {
    return formatter.format(diffInSeconds, 'second');
  }

  const diffInMinutes = Math.floor(diffInSeconds / 60);
  if (diffInMinutes < 60) {
    return formatter.format(diffInMinutes, 'minute');
  }

  const diffInHours = Math.floor(diffInMinutes / 60);
  if (diffInHours < 24) {
    return formatter.format(diffInHours, 'hour');
  }

  const diffInDays = Math.floor(diffInHours / 24);
  return formatter.format(diffInDays, 'day');
}

const expireElement = document.getElementById('expire');
const expireSeconds = document.getElementById('expire').textContent;
expireElement.textContent = relativeTime(expireSeconds, 'en');
