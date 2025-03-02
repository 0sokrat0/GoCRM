import axios from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1';

export interface UserResponse {
  id: string;
  username: string;
  firstName: string;
  lastName: string;
  photoUrl: string;
  Role: string;
  Level: string;
  TelegramID: number;
  ClientName: string;
  LanguageCode: string;
  Phone: string;
  IsPremium: boolean;
  IsBot: boolean;
  LoginDate: string;
  // Добавьте другие поля, если нужно
}

export const getUserByTelegramID = async (tgID: number): Promise<UserResponse> => {
  try {
    const response = await axios.get(`${API_BASE_URL}/users/telegram/${tgID}`);
    return response.data;
  } catch (error) {
    console.error("Ошибка при получении пользователя:", error);
    throw error;
  }
};
