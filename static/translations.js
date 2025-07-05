const translations = {
    en: {
        index_heading: "Create Secret",
        index_input: "Enter your secret:",
        index_note: "Your secret will be encrypted and can only be viewed once.",
        index_create: "Create Secret",

        saved_title: "Ephemera Link - Secret Created",
        saved_heading: "Secret Created",
        saved_note: "It can only be viewed once and will expire in",
        saved_copy: "Copy Link to Clipboard",
        saved_create: "Create Another Secret",

        view_title: "Ephemera Link - View Secret",
        view_heading: "View Secret",
        view_show: "Show Secret",

        retrieve_title: "Ephemera Link - View Secret",
        retrieve_heading: "View Secret",
        retrieve_note: "Your Secret:",
        retrieve_copy: "Copy Secret to Clipboard",
        retrieve_create: "Create Another Secret",

        error_title: "Ephemera Link - Error",
        error_heading: "Error",
        error_create: "Create New Secret",
    },
    ru: {
        index_heading: "Создать секрет",
        index_input: "Введите ваш секрет:",
        index_note: "Ваш секрет будет зашифрован и может быть просмотрен только один раз.",
        index_create: "Создать секрет",

        saved_title: "Ephemera Link - Секрет создан",
        saved_heading: "Секрет создан",
        saved_note: "Он может быть просмотрен один раз и исчезнет через",
        saved_copy: "Скопировать ссылку",
        saved_create: "Создать другой секрет",

        view_title: "Ephemera Link - Просмотреть секрет",
        view_heading: "Просмотреть секрет",
        view_show: "Показать секрет",

        retrieve_title: "Ephemera Link - Просмотреть секрет",
        retrieve_heading: "Просмотреть секрет",
        retrieve_note: "Ваш секрет:",
        retrieve_copy: "Скопировать секрет",
        retrieve_create: "Создать другой секрет",

        error_title: "Ephemera Link - Ошибка",
        error_heading: "Ошибка",
        error_create: "Создать новый секрет",
    }
};

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

function setLanguage(lang) {
    document.documentElement.lang = lang;
    const elements = document.querySelectorAll('[data-i18n]');
    elements.forEach(el => {
        const key = el.getAttribute('data-i18n');
        if (translations[lang] && translations[lang][key]) {
            el.textContent = translations[lang][key];
        }
    });
}

// Detect browser language or use default
const userLang = navigator.language.startsWith('ru') ? 'ru' : 'en';
setLanguage(userLang);
