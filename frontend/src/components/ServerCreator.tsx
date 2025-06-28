import React, { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Slider } from '@/components/ui/slider';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Badge } from '@/components/ui/badge';
import { Upload, Settings, HardDrive, Folder, Power, PowerOff, Play } from 'lucide-react';
import { toast } from '@/hooks/use-toast';
import LoginInterface from './LoginInterface';

interface Server {
  id: string;
  name: string;
  ip: string;
  ram: number;
  status: 'running' | 'stopped' | 'starting' | 'stopping';
  modpack?: string;
  mods: string[];
}

const ServerCreator = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [username, setUsername] = useState('');
  const [servers, setServers] = useState<Server[]>([]);
  const [isCreating, setIsCreating] = useState(false);
  const [serverName, setServerName] = useState('');
  const [ramAmount, setRamAmount] = useState([2]);
  const [modpackFile, setModpackFile] = useState<File | null>(null);
  const [modFiles, setModFiles] = useState<FileList | null>(null);
  const [activeTab, setActiveTab] = useState('modpack');

  const handleLogin = (user: string) => {
    setIsLoggedIn(true);
    setUsername(user);
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    setUsername('');
    setServers([]);
  };

  const handleCreateServer = async () => {
    if (!serverName.trim()) {
      toast({
        title: "Error",
        description: "Please enter a server name",
        variant: "destructive"
      });
      return;
    }

    setIsCreating(true);
    
    // Simulate server creation
    setTimeout(() => {
      const newServer: Server = {
        id: `server_${Date.now()}`,
        name: serverName,
        ip: `${Math.floor(Math.random() * 255)}.${Math.floor(Math.random() * 255)}.${Math.floor(Math.random() * 255)}.${Math.floor(Math.random() * 255)}:25565`,
        ram: ramAmount[0],
        status: 'starting',
        modpack: modpackFile?.name,
        mods: modFiles ? Array.from(modFiles).map(file => file.name) : []
      };

      setServers(prev => [...prev, newServer]);
      
      // Simulate server starting
      setTimeout(() => {
        setServers(prev => prev.map(server => 
          server.id === newServer.id 
            ? { ...server, status: 'running' as const }
            : server
        ));
        
        toast({
          title: "Server Created!",
          description: `${serverName} is now running at ${newServer.ip}`,
        });
      }, 3000);

      setIsCreating(false);
      setServerName('');
      setRamAmount([2]);
      setModpackFile(null);
      setModFiles(null);
    }, 2000);
  };

  const handleServerAction = (serverId: string, action: 'stop' | 'start') => {
    const server = servers.find(s => s.id === serverId);
    if (!server) return;

    const newStatus = action === 'stop' ? 'stopping' : 'starting';
    const finalStatus = action === 'stop' ? 'stopped' : 'running';
    
    setServers(prev => prev.map(s => 
      s.id === serverId ? { ...s, status: newStatus as any } : s
    ));

    // Simulate server action
    setTimeout(() => {
      setServers(prev => prev.map(s => 
        s.id === serverId ? { ...s, status: finalStatus as any } : s
      ));
      
      toast({
        title: `Server ${action === 'stop' ? 'Stopped' : 'Started'}`,
        description: `${server.name} is now ${finalStatus}`,
      });
    }, 2000);
  };

  const handleModpackUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      setModpackFile(file);
      toast({
        title: "Modpack Uploaded",
        description: `${file.name} ready for installation`,
      });
    }
  };

  const handleModsUpload = (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files;
    if (files) {
      setModFiles(files);
      toast({
        title: "Mods Uploaded",
        description: `${files.length} mod(s) ready for installation`,
      });
    }
  };

  if (!isLoggedIn) {
    return (
      <div className="min-h-screen bg-gradient-to-b from-slate-900 to-slate-800 p-6 flex items-center justify-center">
        <div className="max-w-md w-full">
          <div className="text-center space-y-4 mb-8">
            <h1 className="text-4xl font-bold text-green-400 drop-shadow-lg">
              MINECRAFT SERVER CREATOR
            </h1>
            <p className="text-slate-300 text-sm">
              Login to create and manage your servers
            </p>
          </div>
          <LoginInterface
            isLoggedIn={isLoggedIn}
            username={username}
            onLogin={handleLogin}
            onLogout={handleLogout}
          />
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-b from-slate-900 to-slate-800 p-6">
      <div className="max-w-6xl mx-auto space-y-8">
        <div className="flex justify-between items-center">
          <div className="text-center space-y-4">
            <h1 className="text-4xl font-bold text-green-400 drop-shadow-lg">
              MINECRAFT SERVER CREATOR
            </h1>
            <p className="text-slate-300 text-sm">
              Create and manage your custom Minecraft servers
            </p>
          </div>
          <LoginInterface
            isLoggedIn={isLoggedIn}
            username={username}
            onLogin={handleLogin}
            onLogout={handleLogout}
          />
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          {/* Server Creation Panel */}
          <Card className="minecraft-panel border-slate-600">
            <CardHeader>
              <CardTitle className="text-green-400 text-lg flex items-center gap-2">
                <Settings className="h-5 w-5" />
                CREATE NEW SERVER
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-6">
              <div className="space-y-2">
                <Label htmlFor="serverName" className="text-slate-200 text-sm">
                  Server Name
                </Label>
                <Input
                  id="serverName"
                  value={serverName}
                  onChange={(e) => setServerName(e.target.value)}
                  placeholder="My Awesome Server"
                  className="minecraft-input bg-slate-800 border-slate-600 text-slate-100"
                />
              </div>

              <div className="space-y-4">
                <Label className="text-slate-200 text-sm">
                  RAM Allocation: {ramAmount[0]} GB
                </Label>
                <Slider
                  value={ramAmount}
                  onValueChange={setRamAmount}
                  max={6}
                  min={1}
                  step={1}
                  className="w-full"
                />
                <div className="flex justify-between text-xs text-slate-400">
                  <span>1 GB</span>
                  <span>6 GB</span>
                </div>
              </div>

              <Tabs value={activeTab} onValueChange={setActiveTab} className="w-full">
                <TabsList className="grid w-full grid-cols-2 bg-slate-800">
                  <TabsTrigger value="modpack" className="text-xs">MODPACK</TabsTrigger>
                  <TabsTrigger value="mods" className="text-xs">MODS</TabsTrigger>
                </TabsList>
                
                <TabsContent value="modpack" className="space-y-4">
                  <div className="border-2 border-dashed border-slate-600 rounded-lg p-8 text-center space-y-4 bg-slate-800/50">
                    <Upload className="h-12 w-12 text-slate-400 mx-auto" />
                    <div className="space-y-2">
                      <p className="text-slate-300 text-sm">Upload Modpack Server File</p>
                      <p className="text-slate-500 text-xs">ZIP, JAR, or TAR files supported</p>
                    </div>
                    <Input
                      type="file"
                      accept=".zip,.jar,.tar,.tar.gz"
                      onChange={handleModpackUpload}
                      className="minecraft-input bg-slate-700 border-slate-600"
                    />
                    {modpackFile && (
                      <Badge className="bg-green-600 text-black">
                        {modpackFile.name}
                      </Badge>
                    )}
                  </div>
                </TabsContent>
                
                <TabsContent value="mods" className="space-y-4">
                  <div className="border-2 border-dashed border-slate-600 rounded-lg p-8 text-center space-y-4 bg-slate-800/50">
                    <Folder className="h-12 w-12 text-slate-400 mx-auto" />
                    <div className="space-y-2">
                      <p className="text-slate-300 text-sm">Upload Individual Mods</p>
                      <p className="text-slate-500 text-xs">Select multiple JAR files</p>
                    </div>
                    <Input
                      type="file"
                      accept=".jar"
                      multiple
                      onChange={handleModsUpload}
                      className="minecraft-input bg-slate-700 border-slate-600"
                    />
                    {modFiles && (
                      <div className="flex flex-wrap gap-2">
                        {Array.from(modFiles).map((file, index) => (
                          <Badge key={index} className="bg-orange-600 text-black text-xs">
                            {file.name}
                          </Badge>
                        ))}
                      </div>
                    )}
                  </div>
                </TabsContent>
              </Tabs>

              <Button
                onClick={handleCreateServer}
                disabled={isCreating || !serverName.trim()}
                className="w-full minecraft-button bg-green-600 hover:bg-green-700 text-black font-bold py-3"
              >
                {isCreating ? 'CREATING SERVER...' : 'CREATE SERVER'}
              </Button>
            </CardContent>
          </Card>

          {/* Server List Panel */}
          <Card className="minecraft-panel border-slate-600">
            <CardHeader>
              <CardTitle className="text-blue-400 text-lg flex items-center gap-2">
                <HardDrive className="h-5 w-5" />
                YOUR SERVERS
              </CardTitle>
            </CardHeader>
            <CardContent>
              {servers.length === 0 ? (
                <div className="text-center py-12 space-y-4">
                  <HardDrive className="h-16 w-16 text-slate-600 mx-auto" />
                  <p className="text-slate-400 text-sm">No servers created yet</p>
                  <p className="text-slate-500 text-xs">Create your first server to get started</p>
                </div>
              ) : (
                <div className="space-y-4">
                  {servers.map((server) => (
                    <Card key={server.id} className="bg-slate-800 border-slate-600">
                      <CardContent className="p-4 space-y-3">
                        <div className="flex justify-between items-start">
                          <div>
                            <h3 className="font-bold text-slate-200 text-sm">{server.name}</h3>
                            <p className="text-xs text-slate-400">RAM: {server.ram} GB</p>
                          </div>
                          <Badge 
                            className={
                              server.status === 'running' 
                                ? 'bg-green-600 text-black' 
                                : server.status === 'starting'
                                ? 'bg-yellow-600 text-black'
                                : server.status === 'stopping'
                                ? 'bg-orange-600 text-black'
                                : 'bg-red-600 text-white'
                            }
                          >
                            {server.status.toUpperCase()}
                          </Badge>
                        </div>
                        
                        {server.status === 'running' && (
                          <div className="space-y-2">
                            <div className="bg-slate-900 p-2 rounded border border-slate-600">
                              <p className="text-xs text-slate-400">Server IP:</p>
                              <p className="text-sm text-green-400 font-mono">{server.ip}</p>
                            </div>
                          </div>
                        )}
                        
                        {/* Server Control Buttons */}
                        <div className="flex gap-2">
                          {server.status === 'running' ? (
                            <Button
                              onClick={() => handleServerAction(server.id, 'stop')}
                              size="sm"
                              className="minecraft-button bg-red-600 hover:bg-red-700 text-white"
                            >
                              <PowerOff className="h-4 w-4 mr-1" />
                              STOP
                            </Button>
                          ) : server.status === 'stopped' ? (
                            <Button
                              onClick={() => handleServerAction(server.id, 'start')}
                              size="sm"
                              className="minecraft-button bg-green-600 hover:bg-green-700 text-black"
                            >
                              <Play className="h-4 w-4 mr-1" />
                              START
                            </Button>
                          ) : (
                            <Button
                              disabled
                              size="sm"
                              className="minecraft-button bg-gray-600 text-gray-400 cursor-not-allowed"
                            >
                              <Power className="h-4 w-4 mr-1" />
                              {server.status.toUpperCase()}...
                            </Button>
                          )}
                        </div>
                        
                        {server.modpack && (
                          <div className="flex items-center gap-2">
                            <Badge className="bg-purple-600 text-black text-xs">
                              MODPACK: {server.modpack}
                            </Badge>
                          </div>
                        )}
                        
                        {server.mods.length > 0 && (
                          <div className="space-y-1">
                            <p className="text-xs text-slate-400">Mods ({server.mods.length}):</p>
                            <div className="flex flex-wrap gap-1">
                              {server.mods.slice(0, 3).map((mod, index) => (
                                <Badge key={index} className="bg-orange-600 text-black text-xs">
                                  {mod}
                                </Badge>
                              ))}
                              {server.mods.length > 3 && (
                                <Badge className="bg-slate-600 text-slate-300 text-xs">
                                  +{server.mods.length - 3} more
                                </Badge>
                              )}
                            </div>
                          </div>
                        )}
                      </CardContent>
                    </Card>
                  ))}
                </div>
              )}
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
};

export default ServerCreator;
