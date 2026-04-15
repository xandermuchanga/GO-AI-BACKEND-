/**
 * @license
 * SPDX-License-Identifier: Apache-2.0
 */

/**
 * @license
 * SPDX-License-Identifier: Apache-2.0
 */

import { FileCode, Terminal, Layers, Server, ShieldCheck, Zap } from 'lucide-react';

export default function App() {
  return (
    <div className="min-h-screen bg-slate-50 text-slate-900 font-sans p-8">
      <div className="max-w-4xl mx-auto">
        <header className="mb-12 border-b border-slate-200 pb-8">
          <div className="flex items-center gap-3 mb-4">
            <div className="bg-blue-600 p-2 rounded-lg text-white">
              <Server size={32} />
            </div>
            <h1 className="text-4xl font-bold tracking-tight text-slate-900">
              AI Customer Support API <span className="text-blue-600">(Golang)</span>
            </h1>
          </div>
          <p className="text-xl text-slate-600 max-w-2xl">
            A production-ready REST API built with Go and Gin, integrated with Gemini AI for intelligent customer support automation.
          </p>
        </header>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-12">
          <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-200">
            <div className="flex items-center gap-2 mb-4 text-blue-600 font-semibold">
              <Layers size={20} />
              <h2>Clean Architecture</h2>
            </div>
            <p className="text-slate-600 text-sm">
              Separation of concerns using standard Go patterns: Handlers, Services, Models, and Config.
            </p>
          </div>
          <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-200">
            <div className="flex items-center gap-2 mb-4 text-purple-600 font-semibold">
              <Zap size={20} />
              <h2>Gemini AI Integration</h2>
            </div>
            <p className="text-slate-600 text-sm">
              Leverages Gemini 1.5 Flash for real-time intent classification and professional response generation.
            </p>
          </div>
          <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-200">
            <div className="flex items-center gap-2 mb-4 text-green-600 font-semibold">
              <ShieldCheck size={20} />
              <h2>Production Ready</h2>
            </div>
            <p className="text-slate-600 text-sm">
              Includes Dockerfile, structured logging, error handling, and environment-based configuration.
            </p>
          </div>
          <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-200">
            <div className="flex items-center gap-2 mb-4 text-orange-600 font-semibold">
              <Terminal size={20} />
              <h2>Easy Deployment</h2>
            </div>
            <p className="text-slate-600 text-sm">
              Standard Go modules and Docker support make it easy to deploy to any cloud provider.
            </p>
          </div>
        </div>

        <section className="mb-12">
          <h3 className="text-2xl font-bold mb-6 flex items-center gap-2">
            <FileCode className="text-slate-400" />
            Project Structure
          </h3>
          <div className="bg-slate-900 rounded-lg p-6 text-slate-300 font-mono text-sm overflow-x-auto">
            <pre>{`├── main.go             # Entry point & Server initialization
├── go.mod              # Go module definition
├── config/             # Configuration loading (.env)
├── handlers/           # HTTP request handlers (Gin)
├── services/           # AI Service & Gemini integration
├── models/             # Request/Response data structures
├── routes/             # API route definitions
├── .env.example        # Environment variable template
└── Dockerfile          # Containerization for production`}</pre>
          </div>
        </section>

        <section className="mb-12">
          <h3 className="text-2xl font-bold mb-6 flex items-center gap-2">
            <Terminal className="text-slate-400" />
            Quick Start
          </h3>
          <div className="space-y-4">
            <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-200">
              <h4 className="font-semibold mb-2">1. Install Dependencies</h4>
              <code className="bg-slate-100 px-2 py-1 rounded text-pink-600">go mod tidy</code>
            </div>
            <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-200">
              <h4 className="font-semibold mb-2">2. Set Environment Variables</h4>
              <p className="text-sm text-slate-600 mb-2">Copy .env.example to .env and add your Gemini API Key.</p>
              <code className="bg-slate-100 px-2 py-1 rounded text-pink-600">cp .env.example .env</code>
            </div>
            <div className="bg-white p-6 rounded-xl shadow-sm border border-slate-200">
              <h4 className="font-semibold mb-2">3. Run the Server</h4>
              <code className="bg-slate-100 px-2 py-1 rounded text-pink-600">go run main.go</code>
            </div>
          </div>
        </section>

        <footer className="mt-20 pt-8 border-t border-slate-200 text-center text-slate-500 text-sm">
          <p>Generated by AI Studio • Production-Ready Golang Backend</p>
        </footer>
      </div>
    </div>
  );
}

