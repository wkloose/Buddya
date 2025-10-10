# 🌏 Buddya Backend API (AI Cultural Tour)

Backend service for **Buddya**, an AI-powered digital cultural tour guide that helps users explore Indonesian culture, tourism, culinary heritage, and cultural events.  
This backend integrates **Gemini API** for real-time cultural event data and **Google Maps** for interactive place visualization.

---

## ⚙️ Tech Stack
- **Language:** Go (Golang)  
- **Framework:** Gin  
- **ORM:** GORM  
- **Database:** PostgreSQL  
- **AI Integration:** Gemini API (for cultural event generation)  
- **Maps Integration:** Google Maps Embed / Links  

---

## 🎯 Core Features

### 🏙️ 1. City & Category Search
Users can explore cities by selecting a **category**:
- **Budaya:** Local cultural overview.  
- **Wisata:** Cultural tourism spots with map pins.  
- **Kuliner:** Traditional food places with maps.  
- **Event Budaya:** Real-time cultural events using **Gemini API**.

### 🗺️ 2. Static Tourism & Culinary Data
- Stored in the database via **`Place`** model (seeded manually).  
- Shown on interactive Google Maps.  
- Fast response, no AI usage → efficient for repeated calls.

### 🎭 3. Dynamic Event Retrieval (AI Integration)
- Only *Event Budaya* uses Gemini API.  
- The backend fills prompt templates (`PromptTemplate`) dynamically with `{kota}` values.  
- The AI response is cached in `Event` model to save API usage.  

### 🧠 4. Prompt Management & Logging
- `PromptTemplate`: reusable prompt format per category.  
- `PromptLog`: stores every query–response pair for research reproducibility and auditing.  

---

## 🧩 Data Models

| Model | Description |
|--------|--------------|
| **User** | Stores user credentials for backend access. |
| **City** | Contains city info, coordinates, and cultural descriptions. |
| **Place** | Stores static tourism and culinary data (seeded). |
| **Event** | Stores AI-generated cultural event data with map links. |
| **PromptTemplate** | Saves the text structure of AI prompts per category. |
| **PromptLog** | Logs AI request and response data for analysis. |

---

## 🧭 API Routes Overview

All endpoints are under the prefix `/api`.

| Method | Endpoint | Controller | Description |
|---------|-----------|-------------|--------------|
| `GET` | `/api/budaya?city={city}` | `GetBudayaByCity` | Returns cultural description of the specified city. |
| `GET` | `/api/wisata?city={city}` | `GetWisataByCity` | Returns a list of cultural tourism destinations. |
| `GET` | `/api/kuliner?city={city}` | `GetKulinerByCity` | Returns a list of traditional food places. |
| `GET` | `/api/events?city={city}` | `GetEventByCity` | Retrieves upcoming cultural events (AI integrated). |
| `GET` | `/api/places?city={city}` | `GetPlaces` | Returns combined list of culinary and tourism places. |
| `GET` | `/api/all?city={city}` | `GetAllByCity` | Combines all data (culture, tourism, culinary, events) in one response. |

---

## 🧪 Example JSON

### 🏙️ Request Example
**GET** `/api/all?city=Malang`
```json
{
  "city": "Malang"
}
```
### 📚 Response Example
```json
{
  "budaya": "Budaya khas Malang meliputi seni topeng dan tarian tradisional.",
  "wisata": [
    {
      "nama_tempat": "Museum Topeng",
      "alamat": "Jl. Terusan Dieng, Malang",
      "kategori": "wisata"
    }
  ],
  "kuliner": [
    {
      "nama_tempat": "Bakso President",
      "alamat": "Jl. Batanghari No.5",
      "kategori": "kuliner"
    }
  ],
  "events": [
    {
      "nama_event": "Festival Topeng Malang",
      "tanggal": "2025-11-12",
      "lokasi": "Balai Budaya Malang",
      "link": "https://maps.google.com/..."
    }
  ]
}
```
---
🧠 Research Context

This backend is part of a research project focused on AI-Assisted Cultural Exploration Systems.
Main objectives:

Efficient integration of AI to provide real-time cultural events.

Seamless connection between static data (Places) and dynamic data (AI Events).

Traceable, consistent, and auditable AI results through Prompt Management.

Supporting Indonesia’s digital tourism ecosystem by combining AI + Maps Visualization.

---
📁 Project Structure
- /models        → Database models (City, Place, Event, Prompt, User)
- /routes        → API route definitions
- /controllers   → Business logic for each feature
- /config        → Database & environment setup
- /main.go       → Application entry point
