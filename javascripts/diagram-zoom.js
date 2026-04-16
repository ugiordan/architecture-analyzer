// Add fullscreen toggle to mermaid diagrams
document.addEventListener("DOMContentLoaded", function () {
  // Wait for mermaid to render
  var observer = new MutationObserver(function () {
    document.querySelectorAll(".mermaid").forEach(function (el) {
      if (el.querySelector(".diagram-zoom")) return;
      var btn = document.createElement("button");
      btn.className = "diagram-zoom";
      btn.textContent = "\u26F6";
      btn.title = "Toggle fullscreen";
      btn.addEventListener("click", function (e) {
        e.stopPropagation();
        el.classList.toggle("fullscreen");
        btn.textContent = el.classList.contains("fullscreen") ? "\u2715" : "\u26F6";
      });
      el.style.position = "relative";
      el.appendChild(btn);
    });
  });
  observer.observe(document.body, { childList: true, subtree: true });
});

// ESC to exit fullscreen
document.addEventListener("keydown", function (e) {
  if (e.key === "Escape") {
    document.querySelectorAll(".mermaid.fullscreen").forEach(function (el) {
      el.classList.remove("fullscreen");
      var btn = el.querySelector(".diagram-zoom");
      if (btn) btn.textContent = "\u26F6";
    });
  }
});
