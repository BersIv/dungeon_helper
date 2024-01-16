package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import android.widget.SearchView
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.text.Spannable
import android.text.SpannableString
import android.text.style.ForegroundColorSpan
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentInfoAllChaptersBinding

class InfoAllChapters : Fragment() {

    companion object {
        fun newInstance() = InfoAllChapters()
    }

    private lateinit var viewModel: InfoAllChaptersViewModel
    private var _binding: FragmentInfoAllChaptersBinding? = null
    private val binding get() = _binding!!
    private lateinit var gameVersionView: TextView
    private lateinit var gameTitleView: TextView
    private lateinit var gameChaptersView: TextView
    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val infoAllChaptersViewModel = ViewModelProvider(this)[InfoAllChaptersViewModel::class.java]
        _binding = FragmentInfoAllChaptersBinding.inflate(inflater, container,false)
        val root: View = binding.root
        val textView:TextView = binding.textInfoChapter
        infoAllChaptersViewModel.text.observe(viewLifecycleOwner){
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        gameTitleView = binding.gameTitle
        gameTitleView.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_infoAllChapters_to_navigation_info)
        }
        gameVersionView = binding.gameVersion
        gameVersionView.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_infoAllChapters_to_infoVers)
        }
        gameChaptersView = binding.gameChapter1
        gameChaptersView.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_infoAllChapters_to_infoChapter1Content)
        }
        val searchView = binding.searchView
        searchView.setOnQueryTextListener(object : SearchView.OnQueryTextListener {
            override fun onQueryTextSubmit(query: String?): Boolean {
                return false
            }

            override fun onQueryTextChange(newText: String?): Boolean {
                if (newText.isNullOrEmpty()) {
                    return true
                }

                val gameChaptersText = gameChaptersView.text.toString()
                val spannableString = SpannableString(gameChaptersText)

                // Используйте ForegroundColorSpan для выделения текста
                val startIndex = gameChaptersText.indexOf(newText, ignoreCase = true)
                val endIndex = startIndex + newText.length

                if (startIndex != -1) {
                    spannableString.setSpan(
                        ForegroundColorSpan(resources.getColor(R.color.purple_700)),
                        startIndex,
                        endIndex,
                        Spannable.SPAN_EXCLUSIVE_EXCLUSIVE
                    )
                }

                // Установите текст с использованием SpannableString
                gameChaptersView.text = spannableString

                return true
            }
        })

    }


}