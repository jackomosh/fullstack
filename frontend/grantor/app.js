<<<<<<< HEAD
document.addEventListener('DOMContentLoaded', () => {
    const dropzone = document.getElementById('uploadDropzone');
    const fileInput = document.getElementById('csvFileInput');
    const rosterFeedback = document.getElementById('rosterFeedback');
    const poolForm = document.getElementById('poolDeploymentForm');
    const poolFeedback = document.getElementById('poolFeedback');

    // Drag and Drop Logic Handlers
    dropzone.addEventListener('click', () => fileInput.click());

    dropzone.addEventListener('dragover', (e) => {
        e.preventDefault();
        dropzone.classList.add('active-drag');
    });

    dropzone.addEventListener('dragleave', () => {
        dropzone.classList.remove('active-drag');
    });

    dropzone.addEventListener('drop', (e) => {
        e.preventDefault();
        dropzone.classList.remove('active-drag');
        if (e.dataTransfer.files.length) processRosterUpload(e.dataTransfer.files[0]);
    });

    fileInput.addEventListener('change', () => {
        if (fileInput.files.length) processRosterUpload(fileInput.files[0]);
    });

    async function processRosterUpload(file) {
        if (!file.name.endsWith('.csv')) {
            showFeedback(rosterFeedback, 'Validation Error: System ingestion routes process compiled CSV spreadsheets exclusively[cite: 58].', 'error');
            return;
        }

        rosterFeedback.className = 'feedback-msg';
        rosterFeedback.style.display = 'block';
        rosterFeedback.innerHTML = '<div style="display:flex;align-items:center;gap:0.5rem;"><span class="spinner" style="border-top-color:var(--ifc-blue)"></span> Streaming entries to ledger parsing matrix... [cite: 59]</div>';

        const dataForm = new FormData();
        dataForm.append('roster', file);

        try {
            const response = await fetch('/backend/handlers/roster.go', {
                method: 'POST',
                body: dataForm
            });

            if (!response.ok) throw new Error('Ingestion execution failed. Verification registry rejected rows[cite: 59].');

            showFeedback(rosterFeedback, 'Roster processing complete. Student sub-allocation parameters successfully built into contract rules[cite: 59].', 'success');
        } catch (err) {
            showFeedback(rosterFeedback, err.message, 'error');
        }
    }

    // Pool Deployment Request Pipeline
    poolForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        poolFeedback.className = 'feedback-msg';
        poolFeedback.style.display = 'none';

        const budgetValue = document.getElementById('poolBudget').value;

        try {
            const response = await fetch('/backend/handlers/pool.go', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ amount: parseFloat(budgetValue) })
            });

            if (!response.ok) throw new Error('Escrow initialization rejected by vault controller[cite: 56].');

            showFeedback(poolFeedback, 'Vault escrow initialized. Pool asset tokens locked[cite: 56].', 'success');
            poolForm.reset();
        } catch (error) {
            showFeedback(poolFeedback, error.message, 'error');
        }
    });

    function showFeedback(target, text, stylingClass) {
        target.className = `feedback-msg ${stylingClass}`;
        target.textContent = text;
    }
=======
// Donor Dashboard JavaScript

function logout() {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    window.location.href = '../index.html';
}

function createScholarship() {
    const title = document.getElementById('title').value;
    const coverageType = document.getElementById('coverage-type').value;
    const maxAmount = parseFloat(document.getElementById('max-amount').value);
    const slots = parseInt(document.getElementById('slots').value);
    const deadline = document.getElementById('deadline').value;
    const minGpa = parseFloat(document.getElementById('min-gpa').value) || 0;
    const courses = document.getElementById('courses').value.split(',').map(c => c.trim());
    const years = document.getElementById('years').value.split(',').map(y => parseInt(y.trim())).filter(y => !isNaN(y));

    if (!title || !maxAmount || !slots) {
        alert('Please fill in required fields');
        return;
    }

    const token = localStorage.getItem('token');
    
    fetch('/api/donor/scholarships', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + token
        },
        body: JSON.stringify({
            title,
            coverage_type: coverageType,
            max_amount_per_student: maxAmount,
            number_of_slots: slots,
            eligible_courses: courses,
            eligible_years: years,
            min_gpa: minGpa,
            application_end_date: deadline
        })
    })
    .then(res => res.json())
    .then(data => {
        alert('Scholarship created successfully!');
        document.getElementById('create-scholarship-form').reset();
    })
    .catch(err => {
        console.error(err);
        alert('Error creating scholarship (API not available in demo)');
    });
}

function loadScholarships() {
    const token = localStorage.getItem('token');
    
    fetch('/api/donor/scholarships', {
        headers: { 'Authorization': 'Bearer ' + token }
    })
    .then(res => res.json())
    .then(data => {
        // Update scholarships table
        console.log('Scholarships loaded:', data);
    })
    .catch(err => console.error(err));
}

document.addEventListener('DOMContentLoaded', function() {
    loadScholarships();
>>>>>>> e4099c4c7ff8ad76bb436ac4edac9d35c34f4e0a
});