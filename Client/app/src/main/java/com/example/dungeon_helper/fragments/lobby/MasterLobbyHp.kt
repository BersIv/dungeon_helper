package com.example.dungeon_helper.fragments.lobby

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MasterLobbyActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentMasterLobbyHpBinding

class MasterLobbyHp : Fragment() {

    companion object {
        fun newInstance() = MasterLobbyHp()
    }

    private lateinit var viewModel: MasterLobbyHpViewModel
    private var _binding: FragmentMasterLobbyHpBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val masterLobbyHpViewModel = ViewModelProvider(this)[MasterLobbyHpViewModel::class.java]
        _binding = FragmentMasterLobbyHpBinding.inflate(inflater, container, false)
        val textView: TextView = binding.textPlay
        masterLobbyHpViewModel.text.observe(viewLifecycleOwner){
            textView.text = it
        }
        val root: View = binding.root
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        val backBtn = binding.backBtn
        backBtn.setOnClickListener {
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobbyHp_to_masterLobbyActChoose)
        }
        val minusBtn = binding.minusBtn
        val plusBtn = binding.plusBtn
        val valueTextView = binding.value
        var value = 0
        plusBtn.setOnClickListener {
            value++
            valueTextView.text = value.toString()
        }

        minusBtn.setOnClickListener {
            if (value > 0) {
                value--
                valueTextView.text = value.toString()
            }
        }
        val acceptBtn = binding.acceptBtn

    }

}