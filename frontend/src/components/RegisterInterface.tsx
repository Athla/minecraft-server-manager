
import React, { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { UserPlus } from 'lucide-react';
import { toast } from '@/hooks/use-toast';
import { register } from '@/lib/api';

interface RegisterInterfaceProps {
  onRegisterSuccess: () => void;
  onSwitchToLogin: () => void;
}

const RegisterInterface: React.FC<RegisterInterfaceProps> = ({ onRegisterSuccess, onSwitchToLogin }) => {
  const [email, setEmail] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [isRegistering, setIsRegistering] = useState(false);

  const handleRegister = async () => {
    if (!email.trim() || !username.trim() || !password.trim()) {
      toast({
        title: "Error",
        description: "Please fill in all fields",
        variant: "destructive"
      });
      return;
    }

    setIsRegistering(true);
    try {
      await register(email, username, password);
      toast({
        title: "Registration Successful",
        description: "You can now log in with your credentials.",
      });
      onRegisterSuccess();
    } catch (error) {
      toast({
        title: "Registration Failed",
        description: "An error occurred during registration.",
        variant: "destructive"
      });
    } finally {
      setIsRegistering(false);
    }
  };

  return (
    <Card className="minecraft-panel border-slate-600">
      <CardHeader>
        <CardTitle className="text-orange-400 text-lg flex items-center gap-2">
          <UserPlus className="h-5 w-5" />
          CREATE AN ACCOUNT
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="email" className="text-slate-200 text-sm">
            Email
          </Label>
          <Input
            id="email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Enter your email"
            className="minecraft-input bg-slate-800 border-slate-600 text-slate-100"
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="username" className="text-slate-200 text-sm">
            Username
          </Label>
          <Input
            id="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
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
          onClick={handleRegister}
          disabled={isRegistering}
          className="w-full minecraft-button bg-orange-600 hover:bg-orange-700 text-black font-bold py-3"
        >
          {isRegistering ? 'CREATING ACCOUNT...' : 'CREATE ACCOUNT'}
        </Button>
        <Button
          onClick={onSwitchToLogin}
          variant="link"
          className="w-full text-slate-400 hover:text-slate-200"
        >
          Already have an account? Login here.
        </Button>
      </CardContent>
    </Card>
  );
};

export default RegisterInterface;
