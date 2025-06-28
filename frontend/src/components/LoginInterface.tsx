
import React, { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { LogIn, LogOut, User } from 'lucide-react';
import { toast } from '@/hooks/use-toast';

interface LoginInterfaceProps {
  isLoggedIn: boolean;
  username: string;
  onLogin: (username: string) => void;
  onLogout: () => void;
}

const LoginInterface: React.FC<LoginInterfaceProps> = ({
  isLoggedIn,
  username,
  onLogin,
  onLogout
}) => {
  const [loginUsername, setLoginUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isLogging, setIsLogging] = useState(false);

  const handleLogin = async () => {
    if (!loginUsername.trim() || !password.trim()) {
      toast({
        title: "Error",
        description: "Please enter both username and password",
        variant: "destructive"
      });
      return;
    }

    setIsLogging(true);
    // Simulate login process
    setTimeout(() => {
      onLogin(loginUsername);
      setLoginUsername('');
      setPassword('');
      setIsLogging(false);
      toast({
        title: "Login Successful",
        description: `Welcome back, ${loginUsername}!`,
      });
    }, 1500);
  };

  const handleLogout = () => {
    onLogout();
    toast({
      title: "Logged Out",
      description: "See you next time!",
    });
  };

  if (isLoggedIn) {
    return (
      <div className="flex items-center gap-4 p-4 bg-slate-800 rounded-lg border border-slate-600">
        <div className="flex items-center gap-2 text-green-400">
          <User className="h-5 w-5" />
          <span className="font-bold text-sm">{username}</span>
        </div>
        <Button
          onClick={handleLogout}
          variant="outline"
          size="sm"
          className="minecraft-button bg-red-600 hover:bg-red-700 text-white border-red-700"
        >
          <LogOut className="h-4 w-4 mr-2" />
          LOGOUT
        </Button>
      </div>
    );
  }

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
          <Label htmlFor="username" className="text-slate-200 text-sm">
            Username
          </Label>
          <Input
            id="username"
            value={loginUsername}
            onChange={(e) => setLoginUsername(e.target.value)}
            placeholder="Enter your username"
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
      </CardContent>
    </Card>
  );
};

export default LoginInterface;
