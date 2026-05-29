<<<<<<< HEAD
document.addEventListener('DOMContentLoaded', () => {
    const roleOptions = document.querySelectorAll('.role-option');
    const proceedBtn = document.getElementById('proceedToAuthBtn');
    const selectionScreen = document.getElementById('roleSelectionScreen');
    const authScreen = document.getElementById('credentialAuthScreen');
    const resetNodeBtn = document.getElementById('resetNodeViewBtn');
    const fieldsContainer = document.getElementById('formFieldsContainer');
    const identityForm = document.getElementById('identityForm');
    const feedback = document.getElementById('authFeedback');
    const tabs = document.querySelectorAll('.mode-tab');

    let targetRole = null;
    let authMode = 'login'; 

    roleOptions.forEach(opt => {
        opt.addEventListener('click', () => {
            roleOptions.forEach(o => o.classList.remove('active'));
            opt.classList.add('active');
            targetRole = opt.getAttribute('data-role');
            proceedBtn.disabled = false;
        });
    });

    proceedBtn.addEventListener('click', () => {
        if (!targetRole) return;
        sessionStorage.setItem('activeContextRole', targetRole);
        selectionScreen.classList.add('hidden');
        authScreen.classList.remove('hidden');
        renderDynamicLayoutFields();
    });

    resetNodeBtn.addEventListener('click', () => {
        authScreen.classList.add('hidden');
        selectionScreen.classList.remove('hidden');
        feedback.style.display = 'none';
    });

    tabs.forEach(tab => {
        tab.addEventListener('click', () => {
            tabs.forEach(t => t.classList.remove('active'));
            tab.classList.add('active');
            authMode = tab.getAttribute('data-mode');
            feedback.style.display = 'none';
            renderDynamicLayoutFields();
        });
    });

    function renderDynamicLayoutFields() {
        fieldsContainer.innerHTML = '';
        
        if (authMode === 'login') {
            if (targetRole === 'beneficiary') {
                fieldsContainer.innerHTML = `
                    <div class="form-group">
                        <label for="phoneNumber">Registered Student Mobile Link</label>
                        <input type="tel" id="phoneNumber" class="form-input" placeholder="+254 700 000 000" required>
                    </div>`;
            } else {
                fieldsContainer.innerHTML = `
                    <div class="form-group">
                        <label for="email">Corporate Security Email</label>
                        <input type="email" id="email" class="form-input" placeholder="compliance@entity.org" required>
                    </div>
                    <div class="form-group">
                        <label for="password">Account Security Cipher Passkey</label>
                        <input type="password" id="password" class="form-input" placeholder="••••••••" required>
                    </div>`;
            }
        } else {
            // Precise Sign Up Field Validation Architecture
            if (targetRole === 'beneficiary') {
                fieldsContainer.innerHTML = `
                    <div class="form-group">
                        <label for="fullName">Full Official Legal Name</label>
                        <input type="text" id="fullName" class="form-input" placeholder="Jane Doe" required>
                    </div>
                    <div class="form-group">
                        <label for="nationalId">National Identity Card / Registry Number</label>
                        <input type="text" id="nationalId" class="form-input" placeholder="e.g., 34567890" required>
                    </div>
                    <div class="form-group">
                        <label for="admissionNo">Institutional Valid Student ID/Admission Number</label>
                        <input type="text" id="admissionNo" class="form-input" placeholder="e.g., C01-2345-2026" required>
                    </div>
                    <div class="form-group">
                        <label for="phoneNumber">Mobile Number (Bound to Gov ID Registry)</label>
                        <input type="tel" id="phoneNumber" class="form-input" placeholder="+254 700 000 000" required>
                    </div>`;
            } else {
                fieldsContainer.innerHTML = `
                    <div class="form-group">
                        <label for="entityName">Corporate Entity Name</label>
                        <input type="text" id="entityName" class="form-input" placeholder="Legal Registration Title" required>
                    </div>
                    <div class="form-group">
                        <label for="regulatoryId">Regulatory Body Registry Verification License</label>
                        <input type="text" id="regulatoryId" class="form-input" placeholder="e.g., MOE-Bursary-2026" required>
                    </div>
                    <div class="form-group">
                        <label for="email">Primary Corporate Communication Endpoint</label>
                        <input type="email" id="email" class="form-input" placeholder="admin@entity.org" required>
                    </div>
                    <div class="form-group">
                        <label for="password">Access Security Keyphrase</label>
                        <input type="password" id="password" class="form-input" placeholder="Cryptographically strong requirements" required>
                    </div>`;
            }
        }
    }

    identityForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        feedback.className = 'feedback-msg';
        feedback.style.display = 'none';

        const submitBtn = document.getElementById('authSubmitBtn');
        const originalText = submitBtn.innerHTML;
        submitBtn.disabled = true;
        submitBtn.innerHTML = '<span class="spinner"></span>';

        const handlerEndpoint = authMode === 'login' ? '/backend/handlers/auth.go' : '/backend/handlers/signup.go';
        const payload = { role: targetRole, type: authMode };

        try {
            if (targetRole === 'beneficiary') {
                payload.phone = document.getElementById('phoneNumber').value;
                if (authMode === 'signup') {
                    payload.name = document.getElementById('fullName').value;
                    payload.nationalId = document.getElementById('nationalId').value;
                    payload.studentId = document.getElementById('admissionNo').value;
                }
            } else {
                payload.email = document.getElementById('email').value;
                payload.password = document.getElementById('password').value;
                if (authMode === 'signup') {
                    payload.entityName = document.getElementById('entityName').value;
                    payload.regulatoryId = document.getElementById('regulatoryId').value;
                }
            }

            const response = await fetch(handlerEndpoint, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(payload)
            });

            if (!response.ok) throw new Error('Authorization response constraint check failure. Operation rejected.');

            feedback.textContent = authMode === 'login' 
                ? 'Identity authorized. Cryptographic credentials verified. Loading Workspace...' 
                : 'Registration parameters successfully written to global whitelist ledger. You can now execute log in.';
            
            feedback.classList.add('success');

            if (authMode === 'login') {
                setTimeout(() => { window.location.href = `${targetRole}/dashboard.html`; }, 1200);
            } else {
                submitBtn.disabled = false;
                submitBtn.innerHTML = originalText;
                setTimeout(() => { document.getElementById('tabLogin').click(); }, 1500);
            }
        } catch (err) {
            submitBtn.disabled = false;
            submitBtn.innerHTML = originalText;
            feedback.textContent = err.message;
            feedback.classList.add('error');
        }
    });
});
=======
// Shared authentication functions

function saveAuth(token, user) {
    localStorage.setItem('token', token);
    localStorage.setItem('user', JSON.stringify(user));
}

function getToken() {
    return localStorage.getItem('token');
}

function getUser() {
    const user = localStorage.getItem('user');
    return user ? JSON.parse(user) : null;
}

function isAuthenticated() {
    return !!getToken();
}

function logout() {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    window.location.href = '../index.html';
}

async function loginAndRedirect(phone, role) {
    try {
        // Request OTP
        await login(phone, role);
        alert('OTP sent to your phone. Please check.');
        
        // In production, show OTP input modal
        const otp = prompt('Enter OTP:');
        if (otp) {
            const result = await verifyOTP(phone, otp);
            saveAuth(result.token, result.user);
            
            // Redirect based on role
            switch (result.user.role) {
                case 'donor':
                    window.location.href = 'grantor/dashboard.html';
                    break;
                case 'school_admin':
                    window.location.href = 'institution/dashboard.html';
                    break;
                case 'student':
                    window.location.href = 'beneficiary/portal.html';
                    break;
                case 'admin':
                    window.location.href = 'admin/dashboard.html';
                    break;
            }
        }
    } catch (error) {
        alert('Login failed: ' + error.message);
    }
}
>>>>>>> e4099c4c7ff8ad76bb436ac4edac9d35c34f4e0a
