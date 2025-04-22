import axios from 'axios';

const API_URL = 'http://localhost:8080';

export const userService = {
    fetchProgrammers: async () => {
        try {
            const response = await axios.get(`${API_URL}/users`);
            return response.data.users;
        } catch (error) {
            throw error.response?.data || error.message;
        }
    },

    deleteProgrammer: async (id) => {
        try {
            const response = await axios.delete(`${API_URL}/users/delete/${id}`);
            return response.data;
        } catch (error) {
            throw error.response?.data || error.message;
        }
    },

    filterProgrammersBySkill: async (skill) => {
        try {
            const response = await axios.get(`${API_URL}/users/skills/${skill}`);
            return response.data;
        } catch (error) {
            throw error.response?.data || error.message;
        }
    }
}
