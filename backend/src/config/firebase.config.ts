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
    admin.initializeApp({
      credential: admin.credential.applicationDefault(),
      projectId: firebaseConfig.projectId,
    });
  }
  return admin;
};

export const getFirestore = () => {
  const admin = initializeFirebase();
  return admin.firestore();
}; 