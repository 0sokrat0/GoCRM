import axios from "axios";

const API_URL = "http://localhost:8080/api/v1/services";

// Получить список услуг
export const fetchServices = async () => {
  try {
    const response = await axios.get(API_URL);
    return response.data;
  } catch (error) {
    console.error("Ошибка при загрузке услуг:", error);
    return [];
  }
};
