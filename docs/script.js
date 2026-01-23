document.addEventListener('DOMContentLoaded', function () {
  const downloadBtn = document.getElementById('download-btn');
  const tableContainer = document.getElementById('csv-table-container');

  const CALENDAR_TYPES = ['solar', 'lunar', 'birthday_solar', 'birthday_lunar'];

  let tableData = [
    ["Name", "Month", "Day", "Year", "Calendar_Type"],
    ["New Year", "1", "1", "2026", "solar"],
    ["Chinese New Year", "1", "1", "2026", "lunar"],
    ["Women's Day", "3", "8", "2026", "solar"],
    ["Dragon Boat Festival", "5", "5", "2026", "lunar"],
    ["Children's Day", "6", "1", "2026", "solar"],
    ["Mid-Autumn Festival", "8", "15", "2026", "lunar"],
    ["Bruce Lee", "11", "27", "1940", "birthday_solar"],
    ["Bruce Lee", "10", "28", "1940", "birthday_lunar"],
  ];

  // Find column indexes once
  const nameColIndex = tableData[0].indexOf('Name');
  const monthColIndex = tableData[0].indexOf('Month');
  const dayColIndex = tableData[0].indexOf('Day');
  const yearColIndex = tableData[0].indexOf('Year');
  const calendarTypeColIndex = tableData[0].indexOf('Calendar_Type');

  let draggedRowIndex = null;

  function getDaysInMonth(year, month) {
    return new Date(year, month, 0).getDate();
  }

  function renderTable() {
    let tableHtml = '<table>';
    // Header
    tableHtml += '<thead><tr>';
    tableHtml += '<th></th>'; // For drag handle
    tableData[0].forEach((header) => {
      tableHtml += `<th>${header}</th>`;
    });
    tableHtml += '<th>Actions</th>'; // Combined header for buttons
    tableHtml += '</tr></thead>';

    // Body
    tableHtml += '<tbody>';
    tableData.slice(1).forEach((row, rowIndex) => {
      tableHtml += `<tr data-row-index="${rowIndex}" draggable="true">`;
      tableHtml += '<td class="drag-handle">::</td>';

      row.forEach((cellText, colIndex) => {
        let cellContent = '';
        if (colIndex === nameColIndex) {
          cellContent = `<input type="text" value="${cellText}" data-col-index="${colIndex}">`;
        } else if (colIndex === monthColIndex || colIndex === dayColIndex || colIndex === yearColIndex || colIndex === calendarTypeColIndex) {
          let options = '';
          if (colIndex === monthColIndex) {
            for (let i = 1; i <= 12; i++) options += `<option value="${i}" ${i == cellText ? 'selected' : ''}>${i}</option>`;
          } else if (colIndex === dayColIndex) {
            const year = row[yearColIndex] || new Date().getFullYear();
            const month = row[monthColIndex] || 1;
            const daysInMonth = getDaysInMonth(year, month);
            for (let i = 1; i <= daysInMonth; i++) options += `<option value="${i}" ${i == cellText ? 'selected' : ''}>${i}</option>`;
          } else if (colIndex === yearColIndex) {
            const currentYear = new Date().getFullYear();
            for (let i = 1900; i <= 2100; i++) options += `<option value="${i}" ${i == cellText ? 'selected' : ''}>${i}</option>`;
          } else if (colIndex === calendarTypeColIndex) {
            CALENDAR_TYPES.forEach(type => options += `<option value="${type}" ${type === cellText ? 'selected' : ''}>${type}</option>`);
          }
          cellContent = `<select data-col-index="${colIndex}">${options}</select>`;
        } else {
          cellContent = `<div contenteditable="true" data-col-index="${colIndex}">${cellText}</div>`;
        }
        tableHtml += `<td>${cellContent}</td>`;
      });
      tableHtml += `<td class="action-buttons"><button class="add-next-row-btn" data-row-index="${rowIndex}">Add Next</button><button class="delete-row-btn" data-row-index="${rowIndex}">Delete</button></td>`;
      tableHtml += '</tr>';
    });
    tableHtml += '</tbody></table>';
    tableContainer.innerHTML = tableHtml;
    addEventListeners();
  }

  function addEventListeners() {
    tableContainer.querySelectorAll('.delete-row-btn').forEach(btn => btn.addEventListener('click', handleDeleteRow));
    tableContainer.querySelectorAll('.add-next-row-btn').forEach(btn => btn.addEventListener('click', handleAddNextRow));
    tableContainer.querySelectorAll('input, select').forEach(el => el.addEventListener('change', handleCellUpdate));
    tableContainer.querySelectorAll('[contenteditable]').forEach(el => el.addEventListener('blur', handleCellUpdate));
    tableContainer.querySelectorAll('tr[draggable="true"]').forEach(row => {
      row.addEventListener('dragstart', handleDragStart);
      row.addEventListener('dragover', handleDragOver);
      row.addEventListener('dragleave', handleDragLeave);
      row.addEventListener('drop', handleDrop);
    });
  }

  function handleCellUpdate(e) {
    const el = e.target;
    const rowIndex = parseInt(el.closest('tr').dataset.rowIndex, 10);
    const colIndex = parseInt(el.dataset.colIndex, 10);
    const value = el.tagName === 'DIV' ? el.innerText : el.value;
    tableData[rowIndex + 1][colIndex] = value;

    if (colIndex === monthColIndex || colIndex === yearColIndex) {
      renderTable(); // Re-render to update day dropdown
    }
  }

  function handleDeleteRow(e) {
    const rowIndex = parseInt(e.target.dataset.rowIndex, 10);
    tableData.splice(rowIndex + 1, 1);
    renderTable();
  }

  function handleAddNextRow(e) {
    const rowIndex = parseInt(e.target.dataset.rowIndex, 10);
    const numCols = tableData[0] ? tableData[0].length : 1;
    const newRow = Array(numCols).fill('');
    if (nameColIndex !== -1) newRow[nameColIndex] = '';
    if (monthColIndex !== -1) newRow[monthColIndex] = '1';
    if (dayColIndex !== -1) newRow[dayColIndex] = '1';
    if (yearColIndex !== -1) newRow[yearColIndex] = new Date().getFullYear().toString();
    if (calendarTypeColIndex !== -1) newRow[calendarTypeColIndex] = CALENDAR_TYPES[0];
    tableData.splice(rowIndex + 2, 0, newRow); // Insert after the current row
    renderTable();
  }

  function tableToCSV() {
    return tableData.map(row => row.join(',')).join('\n');
  }

  downloadBtn.addEventListener('click', function () {
    const csvContent = tableToCSV();
    const blob = new Blob([csvContent], {type: 'text/csv;charset=utf-8;'});
    const link = document.createElement('a');
    const url = URL.createObjectURL(blob);
    link.setAttribute('href', url);
    link.setAttribute('download', 'csv-to-ical.csv');
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  });

  const getLinkBtn = document.getElementById('get-link-btn');
  if (getLinkBtn) {
    getLinkBtn.addEventListener('click', function () {
      const csvContent = tableToCSV();
      try {
        const base64Content = btoa(csvContent);
        const urlEncodedContent = encodeURIComponent(base64Content);
        const subscriptionLink = `https://csv-to-ical.fantasticmao.cn/remote?base64=${urlEncodedContent}`;
        window.open(subscriptionLink, '_blank');
      } catch (e) {
        console.error("Failed to encode CSV content: ", e);
        alert("Could not generate link due to an encoding error.");
      }
    });
  }

  // Drag and Drop handlers
  function handleDragStart(e) {
    draggedRowIndex = parseInt(e.target.dataset.rowIndex, 10);
    e.target.classList.add('dragging');
  }

  function handleDragOver(e) {
    e.preventDefault();
    const targetRow = e.target.closest('tr');
    if (targetRow) targetRow.classList.add('drag-over');
  }

  function handleDragLeave(e) {
    const targetRow = e.target.closest('tr');
    if (targetRow) targetRow.classList.remove('drag-over');
  }

  function handleDrop(e) {
    e.preventDefault();
    const targetRow = e.target.closest('tr');
    const targetIndex = parseInt(targetRow.dataset.rowIndex, 10);

    document.querySelectorAll('.dragging, .drag-over').forEach(el => el.classList.remove('dragging', 'drag-over'));

    if (draggedRowIndex !== null && draggedRowIndex !== targetIndex) {
      const movedRow = tableData.splice(draggedRowIndex + 1, 1)[0];
      tableData.splice(targetIndex + 1, 0, movedRow);
      renderTable();
    }
    draggedRowIndex = null;
  }

  // Initial render
  renderTable();
});
