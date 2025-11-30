import { useState, useCallback } from 'react';

export const useApi = () => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);

    const callApi = useCallback(async (apiCall, ...args) => {
        setLoading(true);
        setError(null);
        try {
            const response = await apiCall(...args);
            return response.data;
        } catch (err) {
            const errorMessage = err.response?.data?.message || err.message || 'Something went wrong';
            setError(errorMessage);
            throw err;
        } finally {
            setLoading(false);
        }
    }, []);

    return { loading, error, callApi };
};