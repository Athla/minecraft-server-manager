import React from 'react';
import { Button } from '@/components/ui/button';
import { LogOut, User } from 'lucide-react';

interface HeaderProps {
  username: string;
  onLogout: () => void;
}

const Header: React.FC<HeaderProps> = ({ username, onLogout }) => {
  return (
    <header className="w-full max-w-4xl mb-8 p-4 flex justify-between items-center bg-slate-800 rounded-lg border border-slate-600">
      <div className="flex items-center gap-2 text-green-400">
        <User className="h-5 w-5" />
        <span className="font-bold text-sm">{username}</span>
      </div>
      <Button
        onClick={onLogout}
        variant="outline"
        size="sm"
        className="minecraft-button bg-red-600 hover:bg-red-700 text-white border-red-700"
      >
        <LogOut className="h-4 w-4 mr-2" />
        LOGOUT
      </Button>
    </header>
  );
};

export default Header;