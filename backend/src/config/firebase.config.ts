import * as admin from 'firebase-admin';

export const firebaseConfig = {
  apiKey: "AIzaSyDCNcMdXW1njGEEKsWJLcU__58xon1G3dY",
  authDomain: "project-phantom-7f30e.firebaseapp.com",
  projectId: "project-phantom-7f30e",
  storageBucket: "project-phantom-7f30e.firebasestorage.app",
  messagingSenderId: "521096454991",
  appId: "1:521096454991:web:5ea97308a726411c93db3a"
};

// Initialize Firebase Admin SDK
export const initializeFirebase = () => {
  if (!admin.apps.length) {
    try {
      // Try to use service account key file first
      const serviceAccountPath = process.env.GOOGLE_APPLICATION_CREDENTIALS;
      if (serviceAccountPath) {
        admin.initializeApp({
          credential: admin.credential.cert(serviceAccountPath),
          projectId: firebaseConfig.projectId,
        });
      } else {
        // Fallback to application default credentials
        admin.initializeApp({
          credential: admin.credential.applicationDefault(),
          projectId: firebaseConfig.projectId,
        });
      }
    } catch (error) {
      console.error('Firebase initialization error:', error);
      throw new Error('Failed to initialize Firebase. Please check your credentials.');
    }
  }
  return admin;
};

export const getFirestore = () => {
  try {
    const admin = initializeFirebase();
    return admin.firestore();
  } catch (error) {
    console.error('Error getting Firestore:', error);
    throw error;
  }
}; 