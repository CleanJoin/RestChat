// TODO: Опционально импортировать часть настроек через параметры окружения.
// TODO: переименовать в settings

export const MAX_LOGIN_LENGTH = 16;
export const LOGIN_REGEX = /^([a-zA-Z0-9_-]){1,}$/;
export const MAX_PASSWORD_LENGTH = 32;
export const PASSWORD_REGEX = /^([a-zA-Z0-9_-]){1,}$/;
export const MAX_MESSAGE_LENGTH = 1024;
export const MAX_MESSAGES_NUMBER = 100;
export const AUTO_UPDATE_INTERVAL_SEC = 5;