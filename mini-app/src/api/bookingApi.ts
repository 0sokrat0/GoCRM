import axios from 'axios';

const API_BASE_URL = 'https://your-api-domain.com/v1';

export interface Booking {
  id: string;
  date: string;
  service: string;
  master: string;
  status: 'confirmed' | 'pending' | 'completed';
}

export const BookingApi = {
  getUpcomingBookings: async (): Promise<Booking[]> => {
    const response = await axios.get(`${API_BASE_URL}/bookings/upcoming`);
    return response.data;
  },

  cancelBooking: async (bookingId: string) => {
    await axios.delete(`${API_BASE_URL}/bookings/${bookingId}`);
  }
};
