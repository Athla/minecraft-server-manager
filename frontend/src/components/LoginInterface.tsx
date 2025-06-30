import React, { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { LogIn, LogOut, User } from 'lucide-react';
import { toast } from '@/hooks/use-toast';
import { login, logout } from '@/lib/api';

interface LoginInterfaceProps {
	onLogin: (username: string, token: string) => void;
	onSwitchToRegister: () => void; // New prop for switching to register
}

const LoginInterface: React.FC<LoginInterfaceProps> = ({ onLogin, onSwitchToRegister }) => {
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');
	const [isLogging, setIsLogging] = useState(false)

	const handleLogin = async () => {
		if (!email.trim() || !password.trim()) {
			toast({
				title: "Error",
				description: "Please enter both email and password",
				variant: "destructive"
			});
			return;
		}

		setIsLogging(true);
		try {
			const response = await login(email, password);
			onLogin(email, response.data.token);
			setEmail('');
			setPassword('');
			toast({
				title: "Login Successful",
				description: `Welcome back, ${email}!`,
			});
		} catch (error) {
			toast({
				title: "Login Failed",
				description: "Please check your credentials and try again.",
				variant: "destructive"
			});
		} finally {
			setIsLogging(false);
		}
	};



	return (
		<Card className="minecraft-panel border-slate-600">
			<CardHeader>
				<CardTitle className="text-orange-400 text-lg flex items-center gap-2">
					<LogIn className="h-5 w-5" />
					LOGIN TO MINECRAFT SERVER CREATOR
				</CardTitle>
			</CardHeader>
			<CardContent className="space-y-4">
				<div className="space-y-2">
					<Label htmlFor="email" className="text-slate-200 text-sm">
						Email
					</Label>
					<Input
						id="email"
						value={email}
						onChange={(e) => setEmail(e.target.value)}
						placeholder="Enter your email"
						className="minecraft-input bg-slate-800 border-slate-600 text-slate-100"
					/>
				</div>
				<div className="space-y-2">
					<Label htmlFor="password" className="text-slate-200 text-sm">
						Password
					</Label>
					<Input
						id="password"
						type="password"
						value={password}
						onChange={(e) => setPassword(e.target.value)}
						placeholder="Enter your password"
						className="minecraft-input bg-slate-800 border-slate-600 text-slate-100"
					/>
				</div>
				<Button
					onClick={handleLogin}
					disabled={isLogging}
					className="w-full minecraft-button bg-orange-600 hover:bg-orange-700 text-black font-bold py-3"
				>
					{isLogging ? 'LOGGING IN...' : 'LOGIN'}
				</Button>
				<Button
					onClick={onSwitchToRegister} // Call the new prop when this button is clicked
					variant="link" // Use a link-like style
					className="w-full text-slate-400 hover:text-slate-200"
				>
					Create account.
				</Button>
			</CardContent>
		</Card>
	);
};

export default LoginInterface
