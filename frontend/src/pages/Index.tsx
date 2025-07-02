
import React, { useState, useEffect } from 'react';
import LoginInterface from '@/components/LoginInterface';
import RegisterInterface from '@/components/RegisterInterface';
import ServerCreator from '@/components/ServerCreator';
import { Button } from '@/components/ui/button';

const Index = () => {
	const [isLoggedIn, setIsLoggedIn] = useState(false);
	const [username, setUsername] = useState('');
	const [token, setToken] = useState<string | null>(null);
	const [showLogin, setShowLogin] = useState(true); // Default to showing login

	useEffect(() => {
		const storedToken = localStorage.getItem('token');
		const storedUsername = localStorage.getItem('username');
		if (storedToken && storedUsername) {
			// Check if token is expired
			import('@/lib/api').then(({ isTokenExpired }) => {
				if (isTokenExpired(storedToken)) {
					// Token is expired, clear storage
					localStorage.removeItem('token');
					localStorage.removeItem('username');
				} else {
					setIsLoggedIn(true);
					setUsername(storedUsername);
					setToken(storedToken);
				}
			});
		}
	}, []);

	const handleLogin = (newUsername: string, newToken: string) => {
		setIsLoggedIn(true);
		setUsername(newUsername);
		setToken(newToken);
		localStorage.setItem('token', newToken);
		localStorage.setItem('username', newUsername);
	};

	const handleLogout = async () => {
		if (token) {
			try {
				const { logout } = await import('@/lib/api');
				await logout(token);
			} catch (error) {
				console.error('Logout API call failed:', error);
			}
		}
		setIsLoggedIn(false);
		setUsername('');
		setToken(null);
		localStorage.removeItem('token');
		localStorage.removeItem('username');
	};

	const handleRegisterSuccess = () => {
		setShowLogin(true); // After successful registration, show the login interface
	};

	const handleSwitchToRegister = () => {
		setShowLogin(false); // Switch to showing the registration interface
	};

	const handleSwitchToLogin = () => {
		setShowLogin(true); // Switch to showing the login interface
	};

	return (
		<div className="min-h-screen bg-gray-900 flex flex-col items-center justify-center p-4">
			<div className="w-full min-w-md">
				<div className="bg-gradient-to-br from-gray-800 to-gray-700 rounded-2xl shadow-2xl border border-gray-700 p-8">
					{isLoggedIn ? (
						<ServerCreator onLogout={handleLogout} />
					) : (
						<div className="space-y-6">
							{showLogin ? (
								<LoginInterface onLogin={handleLogin} onSwitchToRegister={handleSwitchToRegister} />
							) : (
								<RegisterInterface onRegisterSuccess={handleRegisterSuccess} onSwitchToLogin={handleSwitchToLogin} />
							)}
						</div>
					)}
				</div>
			</div>
		</div>
	);
};

export default Index;
