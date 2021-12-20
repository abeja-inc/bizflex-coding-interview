import { useState } from 'react';
import { Clock } from './Clock';
import { ClockProvider } from './ClockContext';
import { Settings } from './Settings';
import './App.css';

export function App() {
  const [refreshInterval, setRefreshInterval] = useState<number | null>(1000);

  return (
    <div className="App">
      <div className="container">
        <ClockProvider localTimeZone="Asia/Tokyo" refreshInterval={refreshInterval}>
          <Clock title="東京" timeZone="Asia/Tokyo" />
          <Clock title="シンガポール" timeZone="Asia/Singapore" />
          <Clock title="ホノルル" timeZone="Pacific/Honolulu" />
          <Clock title="ロサンゼルス" timeZone="America/Los_Angeles" />
          <Clock title="オークランド" timeZone="Pacific/Auckland" />
          <Settings onIntervalChange={setRefreshInterval} />
        </ClockProvider>
      </div>
    </div>
  );
}
