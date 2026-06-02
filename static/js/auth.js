document.addEventListener('DOMContentLoaded', () => {
    const roleButtons = document.querySelectorAll('#roleSelectionScreen button[data-role]');
    const proceedBtn = document.getElementById('proceedToAuthBtn');
    const selectionScreen = document.getElementById('roleSelectionScreen');
    const authScreen = document.getElementById('credentialAuthScreen');
    const resetNodeBtn = document.getElementById('resetNodeViewBtn');
    const fieldsContainer = document.getElementById('formFieldsContainer');
    const identityForm = document.getElementById('identityForm');
    const feedback = document.getElementById('authFeedback');
    const tabs = document.querySelectorAll('#credentialAuthScreen .flex-1');
    const authSubmitBtn = document.getElementById('authSubmitBtn');

    let targetRole = null;
    let authMode = 'login';

    const activeClasses = ['ring-2', 'ring-cyan-500'];

    roleButtons.forEach(btn => {
        btn.addEventListener('click', () => {
            roleButtons.forEach(b => b.classList.remove(...activeClasses));
            btn.classList.add(...activeClasses);
            targetRole = btn.getAttribute('data-role');
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
        if (feedback) feedback.style.display = 'none';
        roleButtons.forEach(b => b.classList.remove(...activeClasses));
        proceedBtn.disabled = true;
        targetRole = null;
    });

    tabs.forEach(tab => {
        tab.addEventListener('click', () => {
            tabs.forEach(t => {
                t.classList.remove('active', 'border-cyan-500', 'text-white');
                t.classList.add('text-gray-500', 'border-transparent');
            });
            tab.classList.add('active', 'border-cyan-500', 'text-white');
            tab.classList.remove('text-gray-500', 'border-transparent');
            authMode = tab.getAttribute('data-mode');
            authSubmitBtn.textContent = authMode === 'login' ? 'Execute Access Request' : 'Submit Enrollment Application';
            if(feedback) feedback.innerHTML = '';
            renderDynamicLayoutFields();
        });
    });

    function renderDynamicLayoutFields() {
        fieldsContainer.innerHTML = '';
        
        const formGroupClasses = 'mb-4';
        const labelClasses = 'block text-sm font-bold text-gray-300 mb-2';
        const inputClasses = 'w-full bg-gray-700 border border-gray-600 rounded-lg p-3 text-white focus:outline-none focus:ring-2 focus:ring-cyan-500';

        if (authMode === 'login') {
            if (targetRole === 'beneficiary') {
                fieldsContainer.innerHTML = `
                    <div class="${formGroupClasses}">
                        <label for="phoneNumber" class="${labelClasses}">Registered Student Mobile Link</label>
                        <input type="tel" id="phoneNumber" class="${inputClasses}" placeholder="+254 700 000 000" required>
                    </div>`;
            } else {
                fieldsContainer.innerHTML = `
                    <div class="${formGroupClasses}">
                        <label for="email" class="${labelClasses}">Corporate Security Email</label>
                        <input type="email" id="email" class="${inputClasses}" placeholder="compliance@entity.org" required>
                    </div>
                    <div class="${formGroupClasses}">
                        <label for="password" class="${labelClasses}">Account Security Cipher Passkey</label>
                        <input type="password" id="password" class="${inputClasses}" placeholder="••••••••" required>
                    </div>`;
            }
        } else { // signup
            if (targetRole === 'beneficiary') {
                fieldsContainer.innerHTML = `
                    <div class="${formGroupClasses}">
                        <label for="fullName" class="${labelClasses}">Full Official Legal Name</label>
                        <input type="text" id="fullName" class="${inputClasses}" placeholder="Jane Doe" required>
                    </div>
                    <div class="${formGroupClasses}">
                        <label for="nationalId" class="${labelClasses}">National Identity Card / Registry Number</label>
                        <input type="text" id="nationalId" class="${inputClasses}" placeholder="e.g., 34567890" required>
                    </div>
                    <div class="${formGroupClasses}">
                        <label for="admissionNo" class="${labelClasses}">Institutional Valid Student ID/Admission Number</label>
                        <input type="text" id="admissionNo" class="${inputClasses}" placeholder="e.g., C01-2345-2026" required>
                    </div>
                    <div class="${formGroupClasses}">
                        <label for="phoneNumber" class="${labelClasses}">Mobile Number (Bound to Gov ID Registry)</label>
                        <input type="tel" id="phoneNumber" class="${inputClasses}" placeholder="+254 700 000 000" required>
                    </div>`;
            } else { // grantor or institution
                fieldsContainer.innerHTML = `
                    <div class="${formGroupClasses}">
                        <label for="entityName" class="${labelClasses}">Corporate Entity Name</label>
                        <input type="text" id="entityName" class="${inputClasses}" placeholder="Legal Registration Title" required>
                    </div>
                    <div class="${formGroupClasses}">
                        <label for="regulatoryId" class="${labelClasses}">Regulatory Body Registry Verification License</label>
                        <input type="text" id="regulatoryId" class="${inputClasses}" placeholder="e.g., MOE-Bursary-2026" required>
                    </div>
                    <div class="${formGroupClasses}">
                        <label for="email" class="${labelClasses}">Primary Corporate Communication Endpoint</label>
                        <input type="email" id="email" class="${inputClasses}" placeholder="admin@entity.org" required>
                    </div>
                    <div class="${formGroupClasses}">
                        <label for="password" class="${labelClasses}">Access Security Keyphrase</label>
                        <input type="password" id="password" class="${inputClasses}" placeholder="Cryptographically strong requirements" required>
                    </div>`;
            }
        }
    }

    identityForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        
        feedback.innerHTML = '';
        feedback.className = 'mt-4 text-sm';

        const submitBtn = document.getElementById('authSubmitBtn');
        const originalText = submitBtn.innerHTML;
        submitBtn.disabled = true;
        submitBtn.innerHTML = `
            <svg class="animate-spin h-5 w-5 mr-3 inline-block" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            Processing...
        `;

        const payload = { role: targetRole, type: authMode };
        
        // This simplified data gathering assumes IDs on the inputs are unique enough
        Array.from(identityForm.elements).forEach(el => {
            if (el.id && el.value) {
                payload[el.id] = el.value;
            }
        });

        // MOCK API CALL
        await new Promise(resolve => setTimeout(resolve, 1500));
        const isSuccess = Math.random() > 0.2; // Simulate 80% success rate

        if (isSuccess) {
            const successMsg = authMode === 'login' 
                ? 'Identity authorized. Loading secure workspace...' 
                : 'Enrollment successful. You may now log in.';
            
            feedback.textContent = successMsg;
            feedback.className = 'mt-4 text-sm text-green-400 p-3 bg-green-900/50 rounded-lg';

            if (authMode === 'login') {
                setTimeout(() => {
                    const dashboard = targetRole === 'beneficiary' ? 'portal.html' : 'dashboard.html';
                    window.location.href = `${targetRole}/${dashboard}`;
                }, 1200);
            } else {
                submitBtn.disabled = false;
                submitBtn.innerHTML = originalText;
                setTimeout(() => { document.getElementById('tabLogin').click(); }, 1500);
            }
        } else {
            submitBtn.disabled = false;
            submitBtn.innerHTML = originalText;
            feedback.textContent = "Authentication Failed. Please check your credentials or network.";
            feedback.className = 'mt-4 text-sm text-red-400 p-3 bg-red-900/50 rounded-lg';
        }
    });

    // Set initial state for tabs
    document.getElementById('tabLogin').click();
});
