const API_URL = "http://localhost:8080";

const handleResponse = async (response: Response) => {
	if (response.status === 401) {
		localStorage.removeItem("token")
		localStorage.removeItem("username")
		window.location.href = "/"
		throw new Error("Unauthorized")
	}

	if (!response.ok) {
		throw new Error("Request failed")
	}

	return response.json()
}

export const login = async (email: string, password: string) => {
	const response = await fetch(`${API_URL}/auth/login`, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({
			email,
			password,
		}),
	});


	return handleResponse(response)
}

export const register = async (email: string, username: string, password: string) => {
	const response = await fetch(`${API_URL}/auth/register`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({ email, username, password }),
	});

	if (!response.ok) {
		throw new Error('Registration failed');
	}

	return handleResponse(response)
};

export const logout = async (token: string) => {
	const response = await fetch(`${API_URL}/auth/logout`, {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
			Authorization: `Bearer ${token}`,
		},
	});

	if (!response.ok) {
		throw new Error("Failed to logout");
	}

	return response.json()
}
