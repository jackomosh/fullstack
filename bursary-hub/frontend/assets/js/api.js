// Shared API wrapper with JWT header injection and error handling

const API_BASE = 'http://localhost:8080/api';

function getAuthHeaders() {
    const token = localStorage.getItem('token');
    return {
        'Content-Type': 'application/json',
        'Authorization': token ? 'Bearer ' + token : ''
    };
}

async function apiRequest(endpoint, options = {}) {
    const defaults = {
        headers: getAuthHeaders()
    };
    
    const config = Object.assign({}, defaults, options);
    
    try {
        const response = await fetch(API_BASE + endpoint, config);
        
        if (response.status === 401) {
            // Token expired or invalid
            localStorage.removeItem('token');
            localStorage.removeItem('user');
            window.location.href = '../index.html';
            return null;
        }
        
        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.message || 'API request failed');
        }
        
        return await response.json();
    } catch (error) {
        console.error('API Error:', error);
        throw error;
    }
}

function login(phone, role) {
    return apiRequest('/auth/login', {
        method: 'POST',
        body: JSON.stringify({ phone, role })
    });
}

function verifyOTP(phone, otp) {
    return apiRequest('/auth/verify-otp', {
        method: 'POST',
        body: JSON.stringify({ phone, otp })
    });
}

function get(endpoint) {
    return apiRequest(endpoint);
}

function post(endpoint, data) {
    return apiRequest(endpoint, {
        method: 'POST',
        body: JSON.stringify(data)
    });
}

function put(endpoint, data) {
    return apiRequest(endpoint, {
        method: 'PUT',
        body: JSON.stringify(data)
    });
}

function del(endpoint) {
    return apiRequest(endpoint, {
        method: 'DELETE'
    });
}